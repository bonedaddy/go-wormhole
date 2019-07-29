package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

//SendMode wraps an enum style byte
//for tracking the mode, or sub-command
//of a send operation
type SendMode byte

const (
	//SendModeFile sending files, and/or directories
	SendModeFile SendMode = iota
	//SendModeDir sending directories recursively
	SendModeDir
	//SendModeText sending textual (string) data
	SendModeText
)

//ClientSend wraps up the CLI option and maintains
//the connection and progress of sending media
type ClientSend struct {
	*Client

	CodeLength    int
	Verify        bool
	IgnoreBad     bool
	IncludeHidden bool

	//Mode declares how, or what, we are intending
	//to send. Either file, directories, or text
	Mode SendMode
}

//FileDesc is a descriptor for a given
//file to be sent across the wire.
type FileDesc struct {
	//LocalPath is the local, sender
	//path of the file
	LocalPath string

	//Path is the purposed path for
	//the receiver to save the file
	//as when downloaded
	Path string

	//Name is the filename alone without
	//path information
	Name string

	//Size is the on-disk size of the
	//resource
	Size int64

	//Checksum represents the file
	//content checksum for verifying
	//it was downloaded properly.
	//This is a SHA1 hash
	Checksum string
}

//NewClientSend returns a new Send object with
//default fields filled in. The mode
//is set to SendModeFile meaning it's expected
//to send files/directories
func NewClientSend() ClientSend {
	return ClientSend{
		Client: &Client{
			RelayAddress:   defRelay,
			TransitAddress: defTransit,
			AppID:          defAppID,
		},

		CodeLength:    defCodeLen,
		Verify:        false,
		IgnoreBad:     false,
		IncludeHidden: false,
		Mode:          SendModeFile,
	}
}

//newSendFromCLI creates a new send object and fills
//it with the options from the CLI context
func newSendFromCLI(c *cli.Context) ClientSend {
	s := ClientSend{
		Client: &Client{
			RelayAddress:   c.GlobalString("relay"),
			TransitAddress: c.GlobalString("transit"),
			AppID:          c.GlobalString("appid"),

			Code: c.String("code"),
		},
		CodeLength:    int(c.Uint("code-length")),
		Verify:        c.Bool("verify"),
		IgnoreBad:     c.Bool("ignore-bad-files"),
		IncludeHidden: c.Bool("include-hidden"),
	}

	if c.Command.HasName("directory") {
		s.Mode = SendModeDir
	} else if c.Command.HasName("text") {
		s.Mode = SendModeText
	}

	return s
}

//CompileFileManifest takes a slice of path arguments
//and iterates through them to build a manifest
//within the Send object. This will stat check
//each to verify validity and to detect if it is
//a directory. The manifest generated will contain
//local links, to the purposed manifest of file-names
//for the receiving client to use.
//If any errors are encountered they are returned and
//the send object is un-touched
func (s *ClientSend) CompileFileManifest(args []string) error {
	files := make([]FileDesc, 0)
	dirs := make([]string, 0)

	//Read the arguments first to figure out what we need
	//to consume
	for _, fd := range args {
		fi, err := os.Stat(fd)
		if err != nil {
			fmt.Printf("failed to open path '%s'\n", fd)
			if !s.IgnoreBad {
				return errors.New("read path failed")
			}
		}

		if fi.IsDir() {
			if strings.HasPrefix(fi.Name(), ".") && !s.IncludeHidden {
				verbosef("skipping hidden folder '%s'", fi.Name())
				continue
			}

			verbosef("adding directory %s to walk list", fi.Name())
			dirs = append(dirs, filepath.Clean(fd))
		} else {
			absFD, err := filepath.Abs(fd)
			if err != nil {
				fmt.Printf("failed to generate absolute path for file '%s'\n", fd)
				return err
			}

			verbosef("added file to list '%s'", fi.Name())
			files = append(files, FileDesc{
				LocalPath: absFD,
				Path:      filepath.Clean(fd),
				Name:      fi.Name(),
				Size:      fi.Size(),
			})
		}
	}

	//Consume directories into the files portion by walking them recursively
	//at the end we should only have the files slice and can disregard
	//the dirs slice entirely
	for _, dir := range dirs {
		verbosef("walking directory %s for files", dir)
		err := filepath.Walk(dir, func(fd string, fi os.FileInfo, err error) error {
			//If something went wrong, do we care?
			if err != nil {
				if s.IgnoreBad { //Option says we don't, so skip
					return nil
				}
				return err
			}

			if fi.IsDir() == false {
				//Check for hidden
				if !s.IncludeHidden && hasHiddenFolder(fd) {
					verbosef("skipping hidden folder")
					return nil //Skip for hidden
				}

				absFD, err := filepath.Abs(fd)
				if err != nil {
					fmt.Printf("failed to generate absolute path for file '%s'\n", fd)
					if s.IgnoreBad { //Option says we don't, so skip
						return nil
					}
					return err
				}

				verbosef("adding file to list '%s'", fi.Name())
				files = append(files, FileDesc{
					LocalPath: absFD,
					Path:      filepath.Clean(fd),
					Name:      fi.Name(),
					Size:      fi.Size(),
				})
			}

			return nil
		})

		if err != nil && s.IgnoreBad == false {
			return err
		}
	}

	normalizePath(&files)

	s.Manifest = make([]FileDesc, 0)
	for _, fd := range files {
		s.Manifest = append(s.Manifest, fd)
	}

	return nil
}

func hasHiddenFolder(path string) bool {
	directory, _ := filepath.Split(path)
	dirs := strings.Split(directory, string(os.PathSeparator))
	for _, dir := range dirs {
		if strings.HasPrefix(dir, ".") {
			return true
		}
	}
	return false
}

func normalizePath(files *[]FileDesc) {
	//Find the smallest common directory if we can
	smallestDir := make([]string, 0)
	firstHit := false
	for _, fd := range *files {
		dir, _ := filepath.Split(fd.LocalPath)
		subdirs := strings.Split(dir, string(os.PathSeparator))
		if len(subdirs) == 0 {
			verbosef("file path is at root %s", fd.LocalPath)
			smallestDir = smallestDir[:0]
			break //The smallest is just zero or empty
		}

		if firstHit == false {
			smallestDir = subdirs
			firstHit = true
			continue
		}

		//Does the first directory match? If not, the whole tree is different
		var ind int
		var sml string
		for ind, sml = range smallestDir {
			if ind >= len(subdirs) {
				break
			}

			if sml != subdirs[ind] {
				//They don't equal anymore, so this directory is off the list
				break
			}
		}

		if ind == 0 {
			//Got back to zero, so directory is empty as smallest
			smallestDir = smallestDir[:0]
			break
		} else if ind < len(smallestDir)-1 {
			verbosef("trimming smallest %v -> %v", smallestDir, smallestDir[:ind])
			smallestDir = smallestDir[:ind] //Trim down
		}
	}

	smallDir := strings.Join(smallestDir, string(os.PathSeparator))
	verbosef("found smallest common directory '%s'", smallDir)
	//Trim down the paths to the smallest common now
	for i := range *files {
		(*files)[i].Path = (*files)[i].LocalPath[len(smallDir):]

		verbosef("final path: %s -> %s", (*files)[i].LocalPath, (*files)[i].Path)
	}
}

//ComputeChecksums iterates the manifest file in the Send
//object, and calculates the SHA1 Checksum of each file
//within it. Returns an error if something goes wrong
//during the process, and the client is not expected
//to proceed
func (s *ClientSend) ComputeChecksums() error {
	for i, fd := range s.Manifest {
		hash, err := checksumFile(fd.LocalPath)
		if err != nil {
			fmt.Printf("failed to compute checksum for file %s\n", fd.Path)
			return err
		}
		verbosef("checksum computed for '%s' = %s", fd.Path, hash)

		s.Manifest[i].Checksum = hash
	}

	return nil
}

func checksumFile(path string) (string, error) {
	var hash string

	file, err := os.Open(path)
	if err != nil {
		return hash, err
	}
	defer file.Close()

	hasher := sha1.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return hash, err
	}

	hashBytes := hasher.Sum(nil)[:20]
	hash = hex.EncodeToString(hashBytes)

	return hash, nil
}

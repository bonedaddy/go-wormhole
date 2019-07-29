package main

/**
 * Wormhole CLI Apllication
 * ========================
 * TODO: Add SSH key exchange for SSL
 * TODO: Add TOR mode
 **/

import (
	"fmt"
	"os"
	"strings"

	"github.com/chris-pikul/go-wormhole/words"
	"github.com/urfave/cli"
)

const (
	//Version holds the CLI application version
	Version = "0.1.0"

	defRelay   = "locahost:4000"
	defTransit = "locahost:4001"
	defAppID   = "wormhole"
	defCodeLen = 2
)

var app *cli.App
var isVerbose bool

func main() {
	app = cli.NewApp()
	app.Name = "wormhole"
	app.Usage = "send/receive files from one device to another, securely"
	app.Description = ` Create a Magic Wormhole and communicate through it

    Wormholes are created by speaking the same magic CODE in two
    different places at the same time.  Wormholes are secure against
    anyone who doesn't use the same code. 
` //Copied direct from the wormhole source

	app.HelpName = "wormhole"
	app.Version = Version

	app.BashComplete = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "send\nget\nrelay")
	}
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "relay, r",
			Value:  defRelay,
			Usage:  "WebSocket `address` of the relay server to use",
			EnvVar: "WORMHOLE_RELAY",
		},
		cli.StringFlag{
			Name:   "transit, t",
			Value:  defTransit,
			Usage:  "TCP `address` of the transit server to use",
			EnvVar: "WORMHOLE_TRANSIT",
		},
		cli.StringFlag{
			Name:  "appid, a",
			Value: defAppID,
			Usage: "sets the `APPID` to use",
		},
		cli.BoolFlag{
			Name:        "verbose, V",
			Usage:       "output verbose information",
			Destination: &isVerbose,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "send",
			Aliases:   []string{"tx"},
			Usage:     "send a file, or text message, to another device",
			ArgsUsage: "[files/directories]",
			HelpName:  "worhole send",
			Action:    runCommandSend,

			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "code-length, c",
					Value: defCodeLen,
					Usage: "set custom code `length`, or number of password parts",
				},
				cli.BoolFlag{
					Name:  "verify, v",
					Usage: "use extra verification prompt (and wait for approval)",
				},
				cli.StringFlag{
					Name:  "code",
					Usage: "optional `CODE` to use instead of system generated one (dangerous)",
				},
			},

			Subcommands: cli.Commands{
				cli.Command{
					Name:      "file",
					ShortName: "f",
					Aliases:   []string{"files"},
					Usage:     "sends files/directories",
					Action:    runCommandSendFiles,

					Flags: []cli.Flag{
						cli.UintFlag{
							Name:  "code-length, c",
							Value: defCodeLen,
							Usage: "set custom code `length`, or number of password parts",
						},
						cli.BoolFlag{
							Name:  "verify, v",
							Usage: "use extra verification prompt (and wait for approval)",
						},
						cli.StringFlag{
							Name:  "code",
							Usage: "optional `CODE` to use instead of system generated one (dangerous)",
						},

						cli.BoolFlag{
							Name:  "ignore-bad-files, i",
							Usage: "ignore any files that are unable to be read",
						},
						cli.BoolFlag{
							Name:  "include-hidden",
							Usage: "include hidden folders by unix dot prefix",
						},
					},
				},
				cli.Command{
					Name:      "text",
					ShortName: "t",
					Aliases:   []string{"txt", "string"},
					Usage:     "send text/string, use argument '-' to run REPL mode",
					Action:    runCommandSendText,

					Flags: []cli.Flag{
						cli.UintFlag{
							Name:  "code-length, c",
							Value: defCodeLen,
							Usage: "set custom code `length`, or number of password parts",
						},
						cli.BoolFlag{
							Name:  "verify, v",
							Usage: "use extra verification prompt (and wait for approval)",
						},
						cli.StringFlag{
							Name:  "code",
							Usage: "optional `CODE` to use instead of system generated one (dangerous)",
						},
					},
				},
			},
		},

		{
			Name:      "get",
			Aliases:   []string{"rx", "recv", "receive"},
			Usage:     "receive files/directories or text messages from another sender",
			ArgsUsage: "code",
			HelpName:  "wormhole get",
			Action:    runCommandReceive,

			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "only-text, t",
					Usage: "only accept text messages, refuse files",
				},
				cli.BoolFlag{
					Name:  "accept-file, y",
					Usage: "auto-accept any file transfers",
				},

				cli.StringFlag{
					Name:  "output, o",
					Value: "",
					Usage: "specific `PATH` to download to (will overwrite)",
				},
			},
		},
	}

	app.Action = runCommandReceive

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func runCommandSend(c *cli.Context) error {
	cmd := newSendFromCLI(c)

	if !c.Args().Present() {
		fmt.Print("Must supply files, or directories, to send\n\n")
		cli.ShowAppHelpAndExit(c, 1)
		return nil
	}

	cmd.CompileFileManifest(c.Args())

	if err := cmd.ComputeChecksums(); err != nil {
		fmt.Printf("Failed to compute checksum of files!\n")
	}

	codeWords := words.GetRandomWords(cmd.CodeLength)
	code := strings.ToLower(strings.Join(codeWords, "-"))
	fmt.Printf("\nWORMHOLE CODE = %s\n", code)

	return nil
}

func runCommandSendFiles(c *cli.Context) error {
	cmd := newSendFromCLI(c)
	fmt.Printf("%v\n", cmd)

	if c.Args().Present() {
		for i, a := range c.Args() {
			fmt.Printf("argument: %d %q\n", i, a)
		}
	}

	return nil
}

func runCommandSendText(c *cli.Context) error {
	cmd := newSendFromCLI(c)
	fmt.Printf("%v\n", cmd)

	if c.Args().Present() {
		for i, a := range c.Args() {
			fmt.Printf("argument: %d %q\n", i, a)
		}
	}
	return nil
}

func runCommandReceive(c *cli.Context) error {
	fmt.Println("receiving stuff")

	return nil
}

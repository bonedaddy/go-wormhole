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
	"strconv"

	"github.com/urfave/cli"
)

const (
	//Version holds the CLI application version
	Version = "0.1.0"

	defRelay   = "locahost:9000"
	defTransit = "locahost:9001"
)

var app *cli.App

func main() {
	app = cli.NewApp()
	app.Name = "wormhole"
	app.Usage = "send/receive files from one device to another, securely"
	app.HelpName = "wormhole"
	app.Version = Version

	app.BashComplete = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "send\nget\nrelay")
	}
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:      "send",
			Usage:     "send a file, or text message, to another device",
			ArgsUsage: "[files/directories]",
			HelpName:  "worhole send",
			Action:    runCommandSend,

			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "relay, r",
					Value:  defRelay,
					Usage:  "URL address of the relay server to use",
					EnvVar: "WORMHOLE_RELAY",
				},
				cli.StringFlag{
					Name:   "transit, t",
					Value:  defTransit,
					Usage:  "URL address of the transit server to use",
					EnvVar: "WORMHOLE_TRANSIT",
				},
				cli.UintFlag{
					Name:  "code-length, c",
					Value: 2,
					Usage: "set custom code length, or number of password parts",
				},
				cli.BoolFlag{
					Name:  "verify, v",
					Usage: "use extra verification prompt (and wait for approval)",
				},

				cli.StringFlag{
					Name:  "code",
					Usage: "optional code to use instead of system generated one (dangerous)",
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
						cli.BoolFlag{
							Name:  "ignore-bad-files, i",
							Usage: "ignore any files that are unable to be read",
						},
					},
				},
				cli.Command{
					Name:      "directory",
					ShortName: "d",
					Aliases:   []string{"directories"},
					Usage:     "send files recursively from directory",
					Action:    runCommandSendDirectory,

					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "ignore-bad-files, i",
							Usage: "ignore any files that are unable to be read",
						},
					},
				},
				cli.Command{
					Name:      "text",
					ShortName: "t",
					Aliases:   []string{"string"},
					Usage:     "send text/string, use argument '-' to run REPL mode",
					Action:    runCommandSendText,
				},
			},
		},

		{
			Name:      "get",
			Usage:     "receive files/directories or text messages from another sender",
			ArgsUsage: "code",
			HelpName:  "wormhole get",
			Action:    runCommandReceive,

			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "relay, r",
					Value:  defRelay,
					Usage:  "URL address of the relay server to use",
					EnvVar: "WORMHOLE_RELAY",
				},
				cli.StringFlag{
					Name:   "transit, t",
					Value:  defTransit,
					Usage:  "URL address of the transit server to use",
					EnvVar: "WORMHOLE_TRANSIT",
				},

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
					Usage: "specific file name/directory to download to (will overwrite)",
				},
			},
		},

		{
			Name:     "relay",
			Usage:    "act as relay server for other devices to connect through",
			HelpName: "wormhole relay",
			Action:   runCommandRelay,

			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "ports, p",
					Value: &cli.StringSlice{"1334"},
					Usage: "ports to bind to for answering relay requests",
				},

				cli.StringFlag{
					Name:  "log-output, l",
					Value: "",
					Usage: "output file to write log messages to",
				},
				cli.StringFlag{
					Name:  "log-level",
					Value: "INFO",
					Usage: "level of logging messages to write, options are [DEBUG|INFO|WARN|ERROR]",
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
	fmt.Println("sending")

	if c.Args().Present() {
		for i, a := range c.Args() {
			fmt.Printf("argument: %d %q\n", i, a)
		}
	}

	return nil
}

func runCommandSendFiles(c *cli.Context) error {
	fmt.Println("sending file")

	if c.Args().Present() {
		for i, a := range c.Args() {
			fmt.Printf("argument: %d %q\n", i, a)
		}
	}

	return nil
}

func runCommandSendDirectory(c *cli.Context) error {
	fmt.Println("sending directory")

	if c.Args().Present() {
		for i, a := range c.Args() {
			fmt.Printf("argument: %d %q\n", i, a)
		}
	}

	return nil
}

func runCommandSendText(c *cli.Context) error {
	fmt.Println("sending text")

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

func runCommandRelay(c *cli.Context) error {
	fmt.Println("acting as relay")

	//setup the logging infrastructure since we may need it here

	//digest the ports option into actual uint types just to double check the type integrity
	ports := make([]uint, len(c.StringSlice("ports")))
	for i, p := range c.StringSlice("ports") {
		t, err := strconv.ParseUint(p, 10, 32)
		if err != nil {
			fmt.Printf("invalid port number in option 'ports' at index %d, expected a unsigned integer", i)
			os.Exit(1)
		}

		ports[i] = uint(t)
	}

	if len(ports) == 0 {
		return cli.NewExitError("'ports' option cannot be empty", 1)
	}

	return nil
}

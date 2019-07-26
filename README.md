# Magic Wormhole

Safely transfer files, or text, from one device to another (peer-to-peer) without the need for third-party accounts or custom servers. Clients can transfer media between devices (directly if possible) using encryption based on mutually shared passwords. The passwords are called "wormhole codes", and are simple phrases made up of two or more words (and a small channel number). Each client, or side of the connection, must enter the same code in order to safely establish a communication channel. Magic Wormhole does not facilitate this code sharing, so it's up to the sending client to notify the receiver of the code. Because of this fact, the security of the transfer can be of the upmost quality.

## About This Package

Fairly faithful Go implementation of [Magic Wormhole](https://github.com/warner/magic-wormhole) client. Based on the original Python application written by Brian Warner. This version is written in Go, so cross-platform binaries are available without the need for a working Python installation. The API implemented here attempts to be as faithful as possible with it's original Python counterpart, so existing Magic Wormhole relay (rendezvous) servers and transit servers, should work seamlessly. 

## Installation

go-wormhole is made using Go, as such, all the platforms Go supports should be acceptable for go-wormhole as well. Binaries will be provided as I make releases, but "from source" is always an option if you have a Go installation ready.

#### From Source

Requires Go toolkit version 1.12.5 or higher. Due to the dependency on [go-sqlite3](https://github.com/mattn/go-sqlite3) `gcc` is also required, most likely with the `CGO_ENABLED` environment variable as well.

1. `go get github.com/chris-pikul/go-wormhole`
2. `go install github.com/chris-pikul/go-wormhole`

## Usage

go-wormhole is a CLI application (command line), bring up your favourite command line/terminal to start using. Executing `go-wormhole help` will display the help info for usage instructions and available options:

```
COMMANDS:
     get      receive files/directories or text messages from another sender
     send     send a file, or text message, to another device
     help, h  Shows a list of commands or help for one command
```

Running go-wormhole with no command will assume the `get` command for receiving media.

### Receiving Files/Directories/Text

`go-wormhole [wormhole code]` will start the receiving process and start communicating to find the other peer. The wormhole code needs to be the same one the sender initiated with, and they will (or go find them and ask) tell you the code to enter. From here, it should start the transfer if everything is in order proper.

### Sending Files/Directories

`go-wormhole send [file, directories]` will initiate a file transfer. After the files are verified (for existence at least), the wormhole code will print for you to communicate to who you want to receive the files. At this time sending/receiving is only one-to-one at the moment.

### Sending Text

`go-wormhole send text [optional message]` will initiate a transfer of text messages. The wormhole code will print for you to communicate to who you want to receive the messages. The message portion is optional, and if not provided will put the program into REPL mode. In REPL mode you can type out the message line by line until the _end-of-file_ character is entered (`ctrl-d` for linux/mac, `ctrl-z` windows).

## License

This codebase, written by Chris Pikul is licensed under MIT License, see LICENSE for more details.

The original library that this package is based on is also MIT licensed.
Original library: [github.com/warner/magic-wormhole](https://github.com/warner/magic-wormhole)
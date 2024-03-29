package msg

import (
	"encoding/json"
)

//Ack is a server echo response that a
//command message was received
type Ack struct {
	*Message

	ID string `json:"id"`
}
func (m *Ack) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Pong is a protocol specific
//ping/pong test from client to server.
//Pong is sent from server in response
type Pong struct {
	*Message
	Pong int `json:"pong"`
}
func (m *Pong) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//WelcomeInfo wraps up the welcome information
//for servers to respond to clients with.
//This is JSON encoded so the type is provided
//for convenience.
type WelcomeInfo struct {
	//MOTD Message Of The Day. Will be displayed
	//to connecting clients when received
	MOTD *string `json:"motd,omitempty"`

	//Error displays an error message to clients
	//on connection and they will close immediately.
	//Used to anounce service termination
	Error *string `json:"error,omitempty"`

	//Version advertises a new CLI client version to
	//the clients
	Version *string `json:"current_cli_version,omitempty"`
}

//Welcome used for the server to introduce itself to clients
//when they connect
type Welcome struct {
	*Message

	Info WelcomeInfo `json:"welcome"`
}
func (m *Welcome) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Nameplates is the server response to a clients
//list request (List), and responds with the available
//nameplates
type Nameplates struct {
	*Message

	Nameplates []NameplateEntry `json:"nameplates"`
}
func (m *Nameplates) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//NameplateEntry is an individual entry in the
//nameplate response, identified by it's ID
type NameplateEntry struct {
	ID string `json:"id"`
}

//Allocated server response to the Allocate.
//Responds with the allocated nameplate
type Allocated struct {
	*Message

	Nameplate string `json:"nameplate"`
}
func (m *Allocated) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Claimed is the server response to the Claim
//responding with the claimed mailbox
type Claimed struct {
	*Message

	Mailbox string `json:"mailbox"`
}
func (m *Claimed) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Released is the servers confirmation that
//it has released the clients nameplate/mailbox
//as instructed from the Release
type Released struct {
	*Message
}
func (m *Released) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//MailboxMessage is the server pushing a mailbox
//message entry to the client
//The body field is hex encoded.
type MailboxMessage struct {
	*Message

	Side  string `json:"side"`
	Phase string `json:"phase"`
	Body  string `json:"body"` //hex encoded
	MsgID string `json:"msg_id"`
}
func (m *MailboxMessage) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Closed confirms to the client that
//the server has closed their mailbox binding
type Closed struct {
	*Message
}
func (m *Closed) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Error is a server response to errors
//in client messages.
type Error struct {
	*Message

	Error string `json:"error"`
	Orig  Echo   `json:"orig"` //original message
}
func (m *Error) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

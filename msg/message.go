package msg

import (
	"fmt"
	"crypto/rand"
	"encoding/hex"
	"time"
)

//IMessage interface is used for tagging message types.
//Basically just wraps them up with a dummy function
//so reflection can grab them a bit easier
type IMessage interface {
	GetType() Type
	GetID() string
}

//Message base object for communication messages
type Message struct {
	Type Type `json:"type"`

	ID *string `json:"id,omitempty"`
	ServerTime UnixTime `json:"server_tx,omitempty"`
}

//GetType implements the IMessage interface
func (m Message) GetType() Type {
	return m.Type
}

//GetID implements the IMessage interface, returning
//the client's message ID for acknowledgement
func (m Message) GetID() string {
	if m.ID == nil {
		return ""
	}
	return *m.ID
}

//NewMessage returns a new Message object
//with the type filled in
func NewMessage(mt Type) *Message {
	//Generate random message ID
	var id string
	idBytes := make([]byte, 6)
	_, err := rand.Read(idBytes)
	if err == nil {
		//Ignore I suppose?
		id = hex.EncodeToString(idBytes)
	} else {
		//Ignore I suppose?
		id = fmt.Sprintf("%x", time.Now().Unix())
	}

	return &Message{
		Type: mt,
		ID: &id,
	}
}

//NewServerMessage returns a new Message object
//with the type and server time filled in
func NewServerMessage(mt Type) *Message {
	return &Message{
		Type:       mt,
		ServerTime: NewUnixTime(),
	}
}

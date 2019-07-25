package msg

//IMessage interface is used for tagging message types.
//Basically just wraps them up with a dummy function
//so reflection can grab them a bit easier
type IMessage interface {
	GetType() Type
}

//Message base object for communication messages
type Message struct {
	Type Type `json:"type"`

	ServerTime UnixTime `json:"server_tx,omitempty"`
}

//GetType implements the IMessage interface
func (m Message) GetType() Type {
	return m.Type
}

//NewMessage returns a new Message object
//with the type filled in
func NewMessage(mt Type) *Message {
	return &Message{
		Type:       mt,
		ServerTime: NewUnixTime(),
	}
}

//NewServerMessage returns a new Message object
//with the type and server time filled in
func NewServerMessage(mt Type) *Message {
	return &Message{
		Type: mt,
	}
}

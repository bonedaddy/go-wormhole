package msg

import "encoding/json"

//Ping is a protocol specific
//ping/pong test from client to server.
//Ping is sent from client
type Ping struct {
	*Message
	Ping int `json:"ping"`
}

func (m *Ping) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Bind allows the client to declare itself to the server
//and should be the first message the client sends. No other
//client messages will be accepted until this one is received.
type Bind struct {
	*Message

	AppID string `json:"appid"`
	Side  string `json:"side"`
}

func (m *Bind) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//List is a request from the client to have the server
//reply with the available nameplates. The server may deny
//this request. The server fulfills this with the Nameplates
type List struct {
	*Message
}

func (m *List) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Allocate is the client asking for a nameplate
//to be reserved for them to use. The server should
//respond with an Allocated response
type Allocate struct {
	*Message
}

func (m *Allocate) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Claim allows the client to reserve an arbitrary
//nameplate of their choosing (given it is available)
type Claim struct {
	*Message

	Nameplate string `json:"nameplate"`
}

func (m *Claim) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Release tells the server that the client is done
//with their nameplate/mailbox. The Nameplate field
//is optional, but should match the same one from
//Claim, or what is allocated to the client.
type Release struct {
	*Message

	Nameplate string `json:"nameplate"` //optional
}

func (m *Release) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Open binds the client to a mailbox.
//Can only be called once per connection.
//After this point, old messages are transmited
//as well as listening for updates to the mailbox
//in which it will push new ones to the client
type Open struct {
	*Message

	Mailbox string `json:"mailbox"`
}
func (m *Open) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Add tells the server to add a message
//to the mailbox for retrieval later.
//The body portion is hex encoded.
type Add struct {
	*Message

	ID string `json:"id"`
	Phase string `json:"phase"`
	Body  string `json:"body"` //hex encoded
}
func (m *Add) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}

//Close tells the server that the client
//is closing, or terminating it's binding to
//the mailbox. It does not release the mailbox.
type Close struct {
	*Message

	Mood    string `json:"mood"`
	Mailbox string `json:"mailbox"` //optional
}
func (m *Close) parse(src []byte) error {
	err := json.Unmarshal(src, m)
	if err != nil {
		return err
	}
	return nil
}
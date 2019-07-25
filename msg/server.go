package msg

//Ack is a server echo response that a
//command message was received
type Ack struct {
	*Message

	ID int `json:"id"`
}

//Pong is a protocol specific
//ping/pong test from client to server.
//Pong is sent from server in response
type Pong struct {
	*Message
	Pong int `json:"pong"`
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

//Nameplates is the server response to a clients
//list request (List), and responds with the available
//nameplates
type Nameplates struct {
	*Message

	Nameplates []interface{} `json:"nameplates"`
}

//Allocated server response to the Allocate.
//Responds with the allocated nameplate
type Allocated struct {
	*Message

	Nameplate string `json:"nameplate"`
}

//Claimed is the server response to the Claim
//responding with the claimed mailbox
type Claimed struct {
	*Message

	Mailbox string `json:"mailbox"`
}

//Released is the servers confirmation that
//it has released the clients nameplate/mailbox
//as instructed from the Release
type Released struct {
	*Message
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

//Closed confirms to the client that
//the server has closed their mailbox binding
type Closed struct {
	*Message
}

//Error is a server response to errors
//in client messages.
type Error struct {
	*Message

	Error string      `json:"error"`
	Orig  interface{} `json:"orig"` //original message
}

package main

//Client interfaces with the relay server
//and performs the entire protocol for
//sending/receiving media. It is "extended"
//by a ClientSend/ClientReceive type for
//their respective actions.
type Client struct {
	RelayAddress   string
	TransitAddress string
	AppID          string

	//Code is the Wormhole Code generated by the
	//application for the client to relay to the
	//receiver in their own way
	Code string

	//Manifest contains the list of files/directories
	//we are intending to send/receive.
	Manifest []FileDesc
}

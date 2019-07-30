package main

import (
	"fmt"
	"strings"

	"github.com/chris-pikul/go-wormhole/msg"
	"github.com/chris-pikul/go-wormhole/words"
)

func (s *ClientSend) handleReady() error {
	//Request the allocation of a nameplate
	verbosef("allocating nameplate from server")
	s.socketConn.WriteJSON(msg.Allocate{
		Message: msg.NewMessage(msg.TypeAllocate),
	})

	return nil
}

func (s *ClientSend) handleClaimed(mbox string) error {
	verbosef("claimed mailbox from server %s", mbox)

	//Generate code
	codeWords := words.GetRandomWords(s.CodeLength)
	code := strings.ToLower(strings.Join(codeWords, "-"))
	s.Code = s.nameplate+"-"+code
	fmt.Printf("\nWORMHOLE CODE = %s\n", s.Code)
	fmt.Printf("On the other computer run: go-wormhole %s\n", s.Code)

	return nil
}

func (s *ClientSend) handleMessage(mm msg.MailboxMessage) error {
	verbosef("received mailbox message %v", mm.Body)
	/*
	Side  string `json:"side"`
	Phase string `json:"phase"`
	Body  string `json:"body"` //hex encoded
	MsgID string `json:"msg_id"`
	*/
	return nil
}
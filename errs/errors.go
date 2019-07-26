package errs

import "errors"

//ClientError shadows a normal error
//so we can check against internal errors vs.
//client errors. Client errors get written
//direct to the client as is
type ClientError struct {
	message string
}
func (e ClientError) Error() string {
	return e.message
}
//New creates and returns a new ClientError
//with the provided message
func New(msg string) ClientError {
	return ClientError{ msg }
}

//ErrInternal is delivered to the client instead of
//whatever server generated error would be,
//this is to purposely be vague to the client since
//it isn't their problem anyways
var ErrInternal = errors.New("internal server error")

//Common error types, most of these are copied
//from the python version directly in case
//clients reference them directly
var (
	ErrMissingType = New("missing 'type'")
	ErrUnknownType = New("unknown type")
	ErrPing        = New("ping requires 'ping'")

	ErrBindFirst = New("must bind first")
	ErrBound     = New("already bound")
	ErrBindAppID = New("bind requires 'appid'")
	ErrBindSide  = New("bind requires 'side'")

	ErrAlreadyAllocated = New("you already allocated one, don't be greedy")
	ErrAlreadyClaimed = New("only one claim per connection")
	ErrClaimNameplate = New("claim requires 'nameplate'")
	
	ErrNameplateCrowded = New("nameplate crowded")
	ErrReclaimNameplate = New("you cannot re-claim a nameplate that your side previously released")
	ErrMailboxCrowded = New("mailbox crowded")
)

package errs

import "errors"

//Common error types, these are copied
//from the python version directly
var (
	ErrInternal = errors.New("internal server error")

	ErrMissingType = errors.New("missing 'type'")
	ErrUnknownType = errors.New("unknown type")
	ErrPing        = errors.New("ping requires 'ping'")

	ErrBindFirst = errors.New("must bind first")
	ErrBound     = errors.New("already bound")
	ErrBindAppID = errors.New("bind requires 'appid'")
	ErrBindSide  = errors.New("bind requires 'side'")

	ErrAlreadyAllocated = errors.New("you already allocated one, don't be greedy")
	ErrNameplateCrowded = errors.New("nameplate crowded")
	ErrMailboxCrowded = errors.New("mailbox crowded")
)

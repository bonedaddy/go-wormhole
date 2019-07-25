package errs

import "errors"

//Common error types, these are copied
//from the python version directly
var (
	ErrMissingType = errors.New("missing 'type'")
	ErrUnknownType = errors.New("unknown type")
	ErrPing        = errors.New("ping requires 'ping'")

	ErrBindFirst = errors.New("must bind first")
	ErrBound     = errors.New("already bound")
	ErrBindAppID = errors.New("bind requires 'appid'")
	ErrBindSide  = errors.New("bind requires 'side'")
)

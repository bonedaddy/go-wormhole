package msg

import (
	"encoding/json"
	"errors"
)

//ErrUnknown the message type provided does not match
//any known types
var ErrUnknown = errors.New("unknown message type")

//DetectType takes an incoming JSON string as bytes
//and attempts to find the messages type field by
//unmarshalling it into a dummy object.
//Surprisingly, this was faster then RegExp.
func DetectType(src []byte) (Type, error) {
	var hint Message
	err := json.Unmarshal(src, &hint)
	if err != nil {
		return TypeUnknown, err
	}
	return hint.Type, nil
}

//ParseClient attempts to coerce an incoming json message
//into it's respective IMessage implementing Message struct.
//Will return an error if the type cannot be found (ErrUnknown),
//or an error if the message validation is not met
func ParseClient(src []byte) (Type, IMessage, error) {
	mt, err := DetectType(src)
	if err != nil {
		return TypeUnknown, nil, err
	}

	switch mt {
	case TypePing:
		msg := Ping{Message: NewMessage(TypePing)}
		err = msg.parse(src)
		return mt, msg, nil
	case TypeBind:
		msg := Bind{Message: NewMessage(TypeBind)}
		err = msg.parse(src)
		return mt, msg, nil
	case TypeList:
		msg := List{Message: NewMessage(TypeList)}
		err = msg.parse(src)
		return mt, msg, nil
	case TypeAllocate:
		msg := Allocate{Message: NewMessage(TypeAllocate)}
		err = msg.parse(src)
		return mt, msg, nil
	default:
		return mt, *NewMessage(TypeUnknown), ErrUnknown
	}

}

//ParseServer attempts to coerce an incoming json message
//into it's respective IMessage implementing Message struct.
//Will return an error if the type cannot be found (ErrUnknown),
//or an error if the message validation is not met
func ParseServer(src []byte) (Type, IMessage, error) {
	mt, err := DetectType(src)
	if err != nil {
		return TypeUnknown, nil, err
	}

	switch mt {

	}
	return mt, *NewMessage(TypeUnknown), ErrUnknown
}

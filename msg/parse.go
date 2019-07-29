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
	case TypeClaim:
		msg := Claim{Message: NewMessage(TypeClaim)}
		err = msg.parse(src)
		return mt, msg, nil
	case TypeRelease:
		msg := Release{Message: NewMessage(TypeRelease)}
		err = msg.parse(src)
		return mt, msg, nil
	case TypeOpen:
		msg := Open{Message: NewMessage(TypeOpen)}
		err = msg.parse(src)
		return mt, msg, nil
	case TypeAdd:
		msg := Add{Message: NewMessage(TypeAdd)}
		err = msg.parse(src)
		return mt, msg, nil
	case TypeClose:
		msg := Close{Message: NewMessage(TypeClose)}
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
	case TypeAck:
		msg := Ack{Message: NewMessage(TypeAck)}
		err = msg.parse(src)
		return mt, msg, nil	
	case TypePong:
		msg := Pong{Message: NewMessage(TypePong)}
		err = msg.parse(src)
		return mt, msg, nil	
	case TypeWelcome:
		msg := Welcome{Message: NewMessage(TypeWelcome)}
		err = msg.parse(src)
		return mt, msg, nil	
	case TypeNameplates:
		msg := Nameplates{Message: NewMessage(TypeNameplates)}
		err = msg.parse(src)
		return mt, msg, nil
	case TypeAllocated:
		msg := Allocated{Message: NewMessage(TypeAllocated)}
		err = msg.parse(src)
		return mt, msg, nil	
	case TypeClaimed:
		msg := Claimed{Message: NewMessage(TypeClaimed)}
		err = msg.parse(src)
		return mt, msg, nil	
	case TypeReleased:
		msg := Released{Message: NewMessage(TypeReleased)}
		err = msg.parse(src)
		return mt, msg, nil	
	case TypeMessage:
		msg := MailboxMessage{Message: NewMessage(TypeMessage)}
		err = msg.parse(src)
		return mt, msg, nil	
	case TypeClosed:
		msg := Closed{Message: NewMessage(TypeClosed)}
		err = msg.parse(src)
		return mt, msg, nil	
	case TypeError:
		msg := Error{Message: NewMessage(TypeError)}
		err = msg.parse(src)
		return mt, msg, nil	
	}
	return mt, *NewMessage(TypeUnknown), ErrUnknown
}

package msg

//Type enumerates the available message types accepted
type Type int

//Available message type enumerations
const (
	TypeUnknown Type = iota
	TypeAck
	TypePing
	TypePong

	TypeWelcome
	TypeBind
	TypeList
	TypeNameplates
	TypeAllocate
	TypeAllocated
	TypeClaim
	TypeClaimed
	TypeRelease
	TypeReleased
	TypeOpen
	TypeMessage
	TypeAdd
	TypeClose
	TypeClosed

	TypeError //Important! keep this one last
)

var typeMap = map[Type]string{
	TypeUnknown: "unknown",
	TypeAck:     "ack",
	TypePing:    "ping",
	TypePong:    "pong",

	TypeWelcome:    "welcome",
	TypeBind:       "bind",
	TypeList:       "list",
	TypeNameplates: "nameplates",
	TypeAllocate:   "allocate",
	TypeAllocated:  "allocated",
	TypeClaim:      "claim",
	TypeClaimed:    "claimed",
	TypeRelease:    "release",
	TypeReleased:   "released",
	TypeOpen:       "open",
	TypeMessage:    "message",
	TypeAdd:        "add",
	TypeClose:      "close",
	TypeClosed:     "closed",

	TypeError: "error",
}

//TypeToString converts an incoming message type
//to it's string enum equivalent
func TypeToString(mt Type) string {
	str, ok := typeMap[mt]
	if ok {
		return str
	}
	return "unknown"
}

//TypeFromString converts an incoming string
//to it's Type equivalent, or TypeUnknown
//if no matches
func TypeFromString(str string) Type {
	for mt, check := range typeMap {
		if str == check {
			return mt
		}
	}
	return TypeUnknown
}

func (mt Type) String() string {
	return TypeToString(mt)
}

//FromString converts the incoming string into
//the message type it matches, or TypeUnknown
//if no matches
func (mt *Type) FromString(str string) {
	*mt = TypeFromString(str)
}

//MarshalText converts the enum into a string/text representation
func (mt Type) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

//UnmarshalText converts the string representation, back into the enum
func (mt *Type) UnmarshalText(src []byte) error {
	mt.FromString(string(src))
	return nil
}

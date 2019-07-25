package msg

import (
	"encoding/json"
	"strconv"
	"time"
)

//UnixTime wraps a time.Time pointer so that we can
//control the json encoding/decoding of the value.
//The wormhole protocol has it's time values in
//unix seconds since epoch.
type UnixTime struct {
	Time *time.Time
}

//MarshalJSON reads the json
func (t UnixTime) MarshalJSON() ([]byte, error) {
	if t.Time == nil {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time.Unix())
}

//UnmarshalJSON writes the json
func (t *UnixTime) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		t.Time = nil
		return nil
	}

	sec, err := strconv.ParseInt(string(src), 10, 64)
	if err != nil {
		return err
	}

	nix := time.Unix(sec, 0)
	t.Time = &nix
	return nil
}

//NewUnixTime returns a new UnixTime object filled to time.Now()
func NewUnixTime() UnixTime {
	t := time.Now()
	return UnixTime{&t}
}

package msg

//Echo wraps a []byte slice to fool the JSON encoder
//into not string encoding this value. This is used
//in message loopback, or echo-ing the original
//client message while maintaining the JSON format.
type Echo []byte

//MarshalJSON turns this into JSON. For this instance,
//it is a strait copy so no data wrangling happens.
func (e Echo) MarshalJSON() ([]byte, error) {
	return e, nil
}

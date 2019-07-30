package main

import (
	"crypto/rand"
	"encoding/hex"
)

//RandomCode returns a crypto random hexidecimal string
//of given length n
func RandomCode(n int) string {
	var id string
	idBytes := make([]byte, n+1) //note: I got no idea what the perfect "length" is
	_, err := rand.Read(idBytes)
	if err != nil {
		panic("failed to read from random stream: "+err.Error())
	}
	id = hex.EncodeToString(idBytes)

	return id[:n+1]
}
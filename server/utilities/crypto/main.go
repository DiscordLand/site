package crytpo

import (
	"crypto/rand"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var l = uint8(len(chars))

// RandString produces a random string from a predefined list of characters based on length
func RandString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	str := ""
	for _, b := range bytes {
		str += string(chars[b%l])
	}
	return str
}

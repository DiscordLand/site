package main

import (
	"crypto/rand"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randString(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	str := ""
	for _, b := range bytes {
		str += string(chars[b%uint8(len(chars))])
	}
	return str
}

func main() {
	// println(string(bytes))
	// for k, v := range bytes {
	// 	println(k, v)
	// }
	println(randString(38))
}

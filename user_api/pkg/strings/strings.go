package strings

import (
	"math/rand"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "0123456789"
const allBytes = letterBytes + numberBytes

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = allBytes[rand.Intn(len(allBytes))]
	}
	return string(b)
}

func RandNumberBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numberBytes[rand.Intn(len(numberBytes))]
	}
	return string(b)
}

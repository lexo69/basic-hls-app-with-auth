package app

import (
	"math/rand"
	"time"
)

type AccessList map[string]*User

// Create a struct with app users
var AppUsers Users

// Create an access list for the app
var List = make(AccessList)

func RandomString(length int) string {

	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var charset = "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset)-1)]
	}
	return string(b)
}

package utility

import (
	"math/rand"

	"github.com/google/uuid"
)

func CreateUUID() string {
	id := uuid.New()
	result := id.String()
	return result
}

func RandString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678"
	byteString := make([]byte, n)
	for i := range byteString {
		byteString[i] = letters[rand.Intn(len(letters))]
	}
	return string(byteString)

}

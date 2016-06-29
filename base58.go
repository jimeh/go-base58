package base58

import (
	"errors"
	"fmt"
	"strings"
)

// Alphabet is the default alphabet.
const Alphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
const base = len(Alphabet)

// Encode converts a base10 integer to a base58 string using the default
// alphabet.
func Encode(num int) string {
	str := ""
	for num >= base {
		mod := num % base
		str = string(Alphabet[mod]) + str
		num = (num - mod) / base
	}

	return string(Alphabet[num]) + str
}

// Decode converts a base58 string to a base10 integer using the default
// alphabet.
func Decode(str string) (int, error) {
	num := 0
	multi := 1

	for i := len(str); i > 0; i-- {
		char := string(str[i-1])
		index := strings.Index(Alphabet, char)
		if index == -1 {
			return -1, decodeError(str)
		}
		num += multi * index
		multi = multi * base
	}

	return num, nil
}

func decodeError(str string) error {
	msg := fmt.Sprintf("\"%s\" is not a valid base58 string.", str)
	return errors.New(msg)
}

package base58

import (
	"bytes"
	"errors"
)

// Alphabet is the default alphabet.
var Alphabet = []byte(
	"123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ",
)
var base = len(Alphabet)

var errInvalidBase58 = errors.New("invalid base58")

// Encode converts a base10 integer to a base58 string using the default
// alphabet.
func Encode(num int) []byte {
	str := []byte{}

	for num >= base {
		mod := num % base
		str = prepend(str, Alphabet[mod])
		num = (num - mod) / base
	}

	return prepend(str, Alphabet[num])
}

// Decode converts a base58 string to a base10 integer using the default
// alphabet.
func Decode(str []byte) (int, error) {
	num := 0
	multi := 1

	for i := len(str); i > 0; i-- {
		char := str[i-1]
		index := bytes.IndexByte(Alphabet, char)
		if index == -1 {
			return -1, errInvalidBase58
		}
		num += multi * index
		multi = multi * base
	}

	return num, nil
}

func prepend(slice []byte, elem byte) []byte {
	slice = append(slice, byte(0))
	copy(slice[1:], slice)
	slice[0] = elem
	return slice
}

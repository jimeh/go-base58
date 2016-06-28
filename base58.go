package base58

import (
	"errors"
	"strings"
)

// Alphabet is the default alphabet.
const Alphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

// Encode converts a base10 integer to a base58 string using the default
// alphabet.
func Encode(num int) string {
	return EncodeWithAlphabet(num, Alphabet)
}

// Decode converts a base58 string to a base10 integer using the default
// alphabet.
func Decode(str string) (int, error) {
	return DecodeWithAlphabet(str, Alphabet)
}

// EncodeWithAlphabet converts a base10 integer to a base58 string with the
// given alphabet.
func EncodeWithAlphabet(num int, alphabet string) string {
	base := len(alphabet)

	str := ""
	for num >= base {
		mod := num % base
		str = string(alphabet[mod]) + str
		num = (num - mod) / base
	}

	return string(alphabet[num]) + str
}

// DecodeWithAlphabet converts a base58 string to a base10 integer with the
// given alphabet.
func DecodeWithAlphabet(str string, alphabet string) (int, error) {
	base := len(alphabet)
	num := 0
	multi := 1

	for i := len(str); i > 0; i-- {
		char := string(str[i-1])
		index := strings.Index(alphabet, char)
		if index == -1 {
			errorMsg := "\"" + str + "\" is not a valid base58 string."
			return -1, errors.New(errorMsg)
		}
		num += multi * index
		multi = multi * base
	}

	return num, nil
}

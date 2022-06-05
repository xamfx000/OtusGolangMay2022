package hw02unpackstring

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if len(s) == 1 {
		if unicode.IsDigit(rune(s[0])) {
			return "", ErrInvalidString
		}
		return s, nil
	}

	var charToRepeat rune
	result := strings.Builder{}
	for _, char := range s {
		if !unicode.IsDigit(char) {
			if charToRepeat != 0 {
				result.WriteRune(charToRepeat)
			}
			charToRepeat = char
			continue
		}
		if unicode.IsDigit(char) {
			if charToRepeat == 0 {
				return "", ErrInvalidString
			}

			numToRepeat, _ := strconv.Atoi(string(char))
			result.WriteString(
				strings.Repeat(string(charToRepeat), numToRepeat),
			)
			charToRepeat = 0
			continue
		}
	}
	lastChar := rune(s[len(s)-1]) //nolint
	if !unicode.IsDigit(lastChar) {
		result.WriteRune(lastChar)
	}
	return result.String(), nil
}

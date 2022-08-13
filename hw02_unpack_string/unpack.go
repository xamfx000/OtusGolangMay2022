package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
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

	var charToRepeat string
	result := strings.Builder{}
	asSymbolsArray := strings.Split(s, "")
	for _, char := range asSymbolsArray {
		isCharDigit := isDigit(char)
		if !isCharDigit {
			if charToRepeat != "" {
				result.WriteString(charToRepeat)
			}
			charToRepeat = char
			continue
		}
		if isCharDigit {
			if charToRepeat == "" {
				return "", ErrInvalidString
			}

			numToRepeat, _ := strconv.Atoi(char)
			result.WriteString(strings.Repeat(charToRepeat, numToRepeat))
			charToRepeat = ""
			continue
		}
	}
	if lastChar := asSymbolsArray[len(asSymbolsArray)-1]; !isDigit(lastChar) {
		result.WriteString(lastChar)
	}
	return result.String(), nil
}

func isDigit(s string) bool {
	_, conversionError := strconv.Atoi(s)
	return conversionError == nil
}

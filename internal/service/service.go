package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func TextAndMorseCodeConversion(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	isText := strings.ContainsFunc(trimmed, func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	})

	if !isText {
		if !strings.ContainsFunc(trimmed, func(r rune) bool {
			return r == '.' || r == '-' || r == ' '
		}) {
			return "", errors.New("input is ambiguous: not plain text or valid Morse code")
		}
	}

	if isText {
		return morse.ToMorse(trimmed), nil
	}

	return morse.ToText(trimmed), nil
}

package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func TextAndMorseCodeConversion(s string) (string, error) {
	if isMorseCode(s) {
		return morse.ToText(s), nil
	} else if isPlainText(s) {
		return morse.ToMorse(s), nil
	}

	return s, errors.New("the string contains invalid characters")
}

// isMorseCode проверяет, является ли строка кодом Морзе.
func isMorseCode(s string) bool {
	s = strings.TrimSpace(s)
	return !strings.ContainsAny(s, "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ1234567890!?") &&
		strings.ContainsAny(s, ".-")
}

// isPlainText проверяет, является ли строка обычным текстом.
func isPlainText(s string) bool {
	s = strings.TrimSpace(s)
	return !strings.ContainsAny(s, ".-") &&
		strings.ContainsAny(s, "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ1234567890!?")
}

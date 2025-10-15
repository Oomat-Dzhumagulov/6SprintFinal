package service

import (
	"github.com/Oomat-Dzhumagulov/6SprintFinal/pkg/morse"
)

func isMorze(s string) bool {
	for _, char := range s {
		switch char {
		case '.', ',', ' ', '-':
			continue
		default:
			return false
		}
	}
	return true
}

func Convert(s string) string {
	if isMorze(s) {
		return morse.ToText(s)
	}
	return morse.ToMorse(s)
}

package service

import (
	"6sprintFinal/pkg/morse"
)

func isMorze(s string) bool {
	for _, char := range s {
		switch char {
		case '.', ',', ' ', '/':
			continue
		default:
			return false
		}
	}
	return true
}

func Convert(s string) string {
	is := isMorze(s)
	if !is {
		return morse.ToMorse(s)
	}

	return morse.ToText(s)
}

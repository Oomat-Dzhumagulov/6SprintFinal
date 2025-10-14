package service

import "github.com/Yandex-Practicum/go1fl-sprint6-final/morse"

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

package token

import (
	"regexp"
)

func IsAlpha(char string) bool {
	r, _ := regexp.MatchString("[a-z]|[A-Z]", char)
	return r
}

func IsNumeric(char string) bool {
	r, _ := regexp.MatchString("[0-9]", char)
	return r
}

func IsAlphaOrNumeric(char string) bool {
	return IsAlpha(char) || IsNumeric(char)
}

func IsSpace(char string) bool {
	if char == " " {
		return true
	}
	return false
}

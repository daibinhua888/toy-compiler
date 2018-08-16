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

func IsSpecialChar(char string) bool {
	r, _ := regexp.MatchString("_", char)
	return r
}

func IsAlphaOrNumericOrSpecialChar(char string) bool {
	return IsAlpha(char) || IsNumeric(char) || IsSpecialChar(char)
}

func IsSpace(char string) bool {
	if char == " " {
		return true
	}
	return false
}

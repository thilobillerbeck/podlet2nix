package internal

import "unicode"

// Returns a copy of the string s with the first character converted to lowercase.
func LowerFirst(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

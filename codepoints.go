package paasaathai

import (
	"unicode"
)

// Are all the runes in this string in the Thai range of Unicode code points?
func IsThaiString(text string) bool {
	for _, r := range text {
		if !unicode.In(r, unicode.Thai) {
			return false
		}
	}
	return true
}

// Is this rune in the Thai range of Unicode code points?
func IsThaiRune(r rune) bool {
	return unicode.In(r, unicode.Thai)
}

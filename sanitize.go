package busan

import (
	"slices"
	"unicode"
)

const (
	safeStr = "_"
)

func Sanitize(s string) string {
	sans := ""
	for _, ch := range s {
		if slices.Contains(problematicChars, ch) {
			sans += safeStr
		} else if slices.Contains(asciiControlChars, ch) {
			sans += safeStr
		} else if unicode.IsSpace(ch) {
			sans += safeStr
		} else {
			sans += string(ch)
		}
	}

	if slices.Contains(ntfsIntFilenames, sans) {
		sans = safeStr + sans
	}

	return sans
}

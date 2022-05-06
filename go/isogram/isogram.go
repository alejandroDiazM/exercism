package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for _, char := range word {
		if !unicode.IsLetter(char) {
			continue
		} else if charCount := strings.Count(word, string(char)); charCount > 1 {
			return false
		}
	}
	return true
}

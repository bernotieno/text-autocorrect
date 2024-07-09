package punctuation

import (
	"strings"
	"unicode"
)

func Punct(str string) string {
	runes := []rune(str)
	// count := 0
	for i, ch := range runes {
		if unicode.IsPunct(ch) && runes[i-1] == ' ' && ch != '\'' {
			runes[i-1], runes[i] = ch, runes[i-1]
		}
	}

	words := string(runes)
	count := 0
	s := strings.Split(words, " ")
	for i, word := range s {
		// If the word is a single quotation mark and it's the first occurrence, move it to the next word
		if word == "'" && count == 0 {
			count++
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}

	// Adjust punctuation placement for single quotation marks
	for i, word := range s {
		// If the word is a single quotation mark and it's not the first occurrence, move it to the previous word
		if word == "'" {
			s[i-1] = s[i-1] + word
			s = append(s[:i], s[i+1:]...)
		}
	}
	joined := strings.Join(s, " ")

	return joined
}

package transform

import (
	"fmt"
	"strings"
	"unicode"

	"go-reloaded/capitalize"
	"go-reloaded/deconv"
	punctuation "go-reloaded/punctuations"
	"go-reloaded/trim"
	"go-reloaded/vowels"
)

func TransformWords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		switch word {
		case "(hex)":
			// Convert previous word to decimal
			words[i-1] = deconv.ConvToDec(words[i-1], 16)
			// Remove current word from slice
			words = append(words[:i], words[i+1:]...)
		case "(bin)":
			// Convert previous word to decimal
			words[i-1] = deconv.ConvToDec(words[i-1], 2)
			words = append(words[:i], words[i+1:]...)
		case "(up)":
			// Convert previous word to uppercase
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(low)":
			// Convert previous word to lowercase
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(cap)":
			// Capitalize the previous word
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(cap,":
			num := trim.Number(words[i+1])
			if num == 0 {
				fmt.Printf("Error: (cap, %v)is not a valid Number.", num)
				fmt.Println()
				continue
			} else if num < 0 {
				fmt.Printf("Error: (cap, %v) is a negative number.", num)
				fmt.Println()
				continue
			} else if len(words[:i]) < num {
				fmt.Printf("Error: value found in (cap, %v), is greater than %v.\n", num, len(words[:i]))
				for j := 1; j <= len(words[:i]); j++ {
					words[i-j] = capitalize.Capitalized(words[i-j])
				}
			} else {
				for j := 1; j <= num; j++ {
					words[i-j] = capitalize.Capitalized(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
		case "(up,":
			num := trim.Number(words[i+1])
			if num == 0 {
				fmt.Printf("Error: (up, %v)is not a valid Number.", num)
				fmt.Println()
				continue
			} else if num < 0 {
				fmt.Printf("Error: (up, %v) is a negative number.", num)
				fmt.Println()
				continue
			} else if len(words[:i]) < num {
				fmt.Printf("Error: value found in (up, %v), is greater than %v.\n", num, len(words[:i]))
				for j := 1; j <= len(words[:i]); j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			} else {
				for j := 1; j <= num; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)
		case "(low,":
			num := trim.Number(words[i+1])
			if num == 0 {
				fmt.Printf("Error: (low, %v)is not a valid Number.", num)
				fmt.Println()
				continue
			} else if num < 0 {
				fmt.Printf("Error: (low, %v) is a negative number.", num)
				fmt.Println()
				continue
			} else if len(words[:i]) < num {
				fmt.Printf("Error: value found in (low, %v), is greater than %v.\n", num, len(words[:i]))
				for j := 1; j <= len(words[:i]); j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
			} else {
				for j := 1; j <= num; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
			}
			words = append(words[:i], words[i+2:]...)

		case "a", "A":
			if vowels.IsVowel(strings.ToLower(words[i+1])) && word == "a" {
				words[i] += "n"
			} else if vowels.IsVowel(words[i+1]) && word == "A" && unicode.IsLower(rune(words[i+1][0])) {
				words[i] += "n"
			} else if vowels.IsVowel(strings.ToLower(words[i+1])) && word == "A" && unicode.IsUpper(rune(words[i+1][0])) && unicode.IsUpper(rune(words[i+1][1])) {
				words[i] += "N"
			} else if vowels.IsVowel(strings.ToLower(words[i+1])) && word == "A" && unicode.IsUpper(rune(words[i+1][0])) && unicode.IsLower(rune(words[i+1][1])) {
				words[i] += "n"
			}
		
		case "an", "An", "AN":
			if !vowels.IsVowel(strings.ToLower(words[i+1])) && word == "an" {
				words[i] = "a"
			} else if !vowels.IsVowel(strings.ToLower(words[i+1])) && (word == "An" || word == "AN") {
				words[i] = "A"
			}
		}
	}
	joined := strings.Join(words, " ")
	wordPunc := punctuation.Punct(joined)

	return wordPunc
}

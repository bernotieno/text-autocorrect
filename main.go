package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)
// convertToDecimal converts a string representing a number in a specific base to decimal.
func convertToDecimal(str string, baseNO int) string {
	val, err := strconv.ParseInt(str, baseNO, 64)
	if err != nil {
		fmt.Println("ERROR OCCURED CONVERTING", err)
	} else {
		str = strconv.Itoa(int(val))
	}
	return str
}

// number extracts and returns the numeric value from a string as per the base number.
func number(str string) int {
	var val string
	for _, char := range str {
		if unicode.IsDigit(char) {
			val += string(char)
		}
	}
	num, _ :=strconv.Atoi(val)
	return num
}

// beginWithVowel checks and modifies words to begin with "a" or "an" based on the next word.
func beginWithVowel(words []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}

	for i, word := range words {
		for _, b := range vowels {
			if word == "a" && string(words[i+1][0]) == b {
				words[i] = "an"
			} else if word == "A" && string(words[i+1][0]) == b {
				words[i] = "An"
			}
		}
	}
	return words
}

// punctuations handles punctuation marks in a slice of strings.
func punctuations(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}

	// Punc at the beginning of a string
	for i, word := range s {
		for _, punc := range puncs {
			if strings.HasPrefix(word, punc) && !strings.HasSuffix(word, punc) {
				s[i-1] += punc
				s[i] = word[1:]
			}
		}
	}

	// Punc at the end of a string
	for i, word := range s {
		for _, punc := range puncs {
			if strings.HasPrefix(word, punc) && s[len(s)-1] == word {
				s[i-1] += word
				s = s[:len(s)-1]
			}
		}
	}

	// Punc in the middle of a string
	for i, word := range s {
		for _, punc := range puncs {
			if strings.HasPrefix(word, punc) && strings.HasSuffix(word, punc) && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	// Handle apostrophes
	count := 0
	for i, word := range s {
		if word == "'" && count == 0 {
			count++
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}

	// Handle second apostrophes
	for i, word := range s {
		if word == "'" {
			s[i-1] = s[i-1] + word
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}


// Wordschanger applies various transformations to a slice of words based on specified rules.
func Wordschanger(words []string) []string {
	for i, word := range words {
		switch word {
			// convert hexadecimal to decimal
		case "(hex)":
			words[i-1] = convertToDecimal(words[i-1], 16)
			words = append(words[:i], words[i+1:]...)
			// convert binary to decimal
		case "(bin)":
			words[i-1] = convertToDecimal(words[i-1], 2)
			words = append(words[:i], words[i+1:]...)
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(cap)":
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)

			// Apply transformations to multiple words.
		case "(cap,":
			val := number(words[i+1])
			for j := 1; j <= val; j++ {
				words[i-j] = strings.Title(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		case "(up,":
			val := number(words[i+1])
			for j := 1; j <= val; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		case "(low,":
			val := number(words[i+1])
			for j := 1; j <= val; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	// apply additonal transformations
	word := beginWithVowel(words)
	result := punctuations(word)
	return result
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]

	// Open the input text file
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var lines string
	// read each line of the file combining them into single string
	for scanner.Scan() {
		lines += scanner.Text() + " "
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}

	// Split the lines into words
	result := strings.Fields(lines)

	// Apply transformations to the words
	finalResult := Wordschanger(result)

	// Convert the final result to a string
	finalStr := strings.Join(finalResult, " ")

	// Write the final result to an output text file
	doc := os.WriteFile(args[1], []byte(finalStr), 0o644)
	if doc != nil {
		panic(doc)
	}
}

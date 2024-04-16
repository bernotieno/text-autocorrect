package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Converts a string representation of a number in a given base to its decimal equivalent.
func convertToDecimal(str string, baseNo int) string {
	// Parse the string as an integer with the given base
	val, err := strconv.ParseInt(str, baseNo, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		// Convert the parsed value to string
		str = strconv.Itoa(int(val))
	}
	return str
}

// Extracts the numerical value from a string and converts it to an integer.
func numberAfterCom(str string) int {
	// Initialize an empty string to store the numerical value
	var val string

	for _, char := range str {
		// Check if the character is a digit, append it to val
		if unicode.IsDigit(char) {
			val += string(char)
		}
	}

	result, _ := strconv.Atoi(val)

	return result
}

// Checks if a given character is a vowel.
func vowelCase(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	default:
		return false
	}
}



// Adjusts the punctuation placement in a slice of words.

func punctuations(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}

	// Adjust punctuation placement in words
	for i, word := range s {
		for _, punc := range puncs {
			// If word starts with punctuation but does not end with it, move the punctuation to the previous word
			if strings.HasPrefix(word, punc) && !strings.HasSuffix(word, punc) {
				s[i-1] += punc
				s[i] = word[1:]
			}
		}
	}

	// Adjust punctuation placement in words for special cases
	for i, word := range s {
		for _, punc := range puncs {
			// If word starts with punctuation and it's the last word, move the punctuation to the previous word
			if strings.HasPrefix(word, punc) && s[len(s)-1] == word {
				s[i-1] += word
				s = s[:len(s)-1]
			}
		}
	}

	// Adjust punctuation placement in words for special cases
	for i, word := range s {
		for _, punc := range puncs {
			// If word starts and ends with punctuation and it's not the last word, move the punctuation to the previous word
			if strings.HasPrefix(word, punc) && strings.HasSuffix(word, punc) && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	// Adjust punctuation placement for single quotation marks
	count := 0
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

	// Return the slice of words with corrected punctuation placement
	return s
}

func wordsChanger(str []string) []string {
	for i, word := range str {
		switch word {
		case "(hex)":
			// Convert hex to decimal
			str[i-1] = convertToDecimal(str[i-1], 16)
			// Remove current word from slice
			str = append(str[:i], str[i+1:]...)
		case "(bin)":
			// Convert binary to decimal
			str[i-1] = convertToDecimal(str[i-1], 2)
			str = append(str[:i], str[i+1:]...)
		case "(up)":
			// Convert previous word to uppercase
			str[i-1] = strings.ToUpper(str[i-1])
			str = append(str[:i], str[i+1:]...)
		case "(low)":
			// Convert previous word to lowercase
			str[i-1] = strings.ToLower(str[i-1])
			str = append(str[:i], str[i+1:]...)
		case "(cap)":
			// Capitalize the previous word
			str[i-1] = strings.Title(str[i-1])
			str = append(str[:i], str[i+1:]...)
		case "(cap,":
			// Capitalize preceding words based on the value after the comma
			val := numberAfterCom(str[i+1])
			for j := 1; j <= val; j++ {
				str[i-j] = strings.Title(str[i-j])
			}
			str = append(str[:i], str[i+2:]...)
		case "(up,":
			// Convert preceding words to uppercase based on the value after the comma
			val := numberAfterCom(str[i+1])
			for j := 1; j <= val; j++ {
				str[i-j] = strings.ToUpper(str[i-j])
			}
			str = append(str[:i], str[i+2:]...)
		case "(low,":
			// Convert preceding words to lowercase based on the value after the comma
			val := numberAfterCom(str[i+1])
			for j := 1; j <= val; j++ {
				str[i-j] = strings.ToLower(str[i-j])
			}
			str = append(str[:i], str[i+2:]...)

		case "a", "A":
			// Append 'n' to 'a' or 'A' if the next word starts with a vowel
			if vowelCase(str[i+1][0]) {
				str[i] += "n"
			}
		}
	}

	// Apply punctuation transformation and return the result
	result := punctuations(str)
	return result
}

func main() {
	// Check if the correct number of command line arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Invalid parameters, usage: go run . par1 par2")
		return
	}

	// Extract command line arguments
	args := os.Args[1:]

	// Open the input file
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err) // Print error if opening file fails
	}
	defer file.Close() // Ensure file is closed after function returns

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var lines string
	// Read lines from file and concatenate them into a single string
	for scanner.Scan() {
		lines += scanner.Text() + " "
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}

	// Split the string into words
	result := strings.Fields(lines)

	// Transform each word according to some logic
	finalResult := wordsChanger(result)

	// Join the transformed words into a single string separated by spaces
	finalStr := strings.Join(finalResult, " ")

	// Write the final string to the output file
	doc := os.WriteFile(args[1], []byte(finalStr), 0o644)
	if doc != nil {
		panic(doc) // Panic if writing to file fails
	}

}

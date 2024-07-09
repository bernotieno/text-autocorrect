package capitalize

func Capitalized(s string) string {
	capitalized := []rune(s)
	wordStart := true
	for i, char := range capitalized {
		if isAlphanumeric(char) {
			if wordStart {
				capitalized[i] = toUpper(char)
				wordStart = false
			} else {
				capitalized[i] = toLower(char)
			}
		} else {
			wordStart = true
		}
	}
	return string(capitalized)
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toUpper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - ('a' - 'A')
	}
	return char
}

func toLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

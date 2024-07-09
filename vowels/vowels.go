package vowels

func IsVowel(str string) bool {
	if str[0] == 'a' || str[0] == 'e' || str[0] == 'i' || str[0] == 'o' || str[0] == 'u' {
		return true
	} else if str == "hour" || str == "hour's" || str == "honor" || str == "heir" || str == "honesty" || str == "honorarium" || str == "herb" || str == "hors" {
		return true
	} else {
		return false
	}
}

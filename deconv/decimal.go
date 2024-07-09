package deconv

import (
	"fmt"
	"strconv"
)

// Converts a string representation of a number in a given base to its decimal equivalent.
func ConvToDec(str string, baseNum int) string {
	// Parse the string as an integer with the given base
	val, err := strconv.ParseInt(str, baseNum, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		// Convert the parsed value to string
		str = strconv.Itoa(int(val))
	}
	return str
}

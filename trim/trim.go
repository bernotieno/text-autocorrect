package trim

import (
	"fmt"
	"strconv"
)

func Number(str string) int {
	num, err := strconv.Atoi(str[:len(str)-1])
	if err != nil {
		fmt.Println("Couldn't obtain the number:", err)
	}
	return num
}

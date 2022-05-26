package luhn

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func Valid(id string) bool {
	//Remove whitespaces
	idString := strings.ReplaceAll(id, " ", "")

	//Check if length is greater than 1
	if len(idString) <= 1 {
		return false
	}

	//Create a slice of digits to perform the luhn operations,
	// exiting if there are non-digit elements
	var err error
	idSlice := make([]int, utf8.RuneCountInString(idString))

	for i, ltr := range []rune(idString) {
		idSlice[i], err = strconv.Atoi(string(ltr))
		if err != nil {
			return false
		}
	}

	//Double every second digit from the right
	for i := len(idSlice) - 2; i >= 0; i -= 2 {
		idSlice[i] *= 2
		if idSlice[i] > 9 {
			idSlice[i] -= 9
		}
	}

	//Sums all the digits in the slice and checks for module 10
	var sum int
	for _, val := range idSlice {
		sum += val
	}
	return sum%10 == 0
}

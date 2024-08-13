package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

func HexToDecimal(input string) string {
	//split the input string to individual words
	words := strings.Fields(input)

	// iterate through words slice to find hex and preceding word

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {

			hexNumber := words[i-1]

			//convert hexNum to decimal
			decimalNumber, _ := strconv.ParseInt(hexNumber, 16, 64)

			words[i-1] = fmt.Sprintf("%d", decimalNumber)

			//remove (hex) from the list of words
			words = append(words[:i], words[i+1:]...)
		}

	}

	modifiedText := strings.Join(words, " ")

	return modifiedText

}

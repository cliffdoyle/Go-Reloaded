package utilities

import (
	"regexp"
	"strconv"
	"strings"
)

func BinarytoDecimal(text string) string {
	// compile Regular expression pattern to match binary numbers (\b matches word boundaries, [01]+ matches one or more binary digits)
	re := regexp.MustCompile(`(\b[01]+\b)\s*\(bin\)?`)

	// Find all binary numbers in the text
	binaryMatches := re.FindAllStringSubmatch(text, -1)

	// Iterate over each binary number found in the text
	for _, match := range binaryMatches {
		// Extract the binary number
		bin := match[1] 
		// Convert the binary number to decimal
		decimal, err := strconv.ParseInt(bin, 2, 64)
		if err != nil {
		
			continue
		}
		// Replace the binary number and accompanying "(bin)" word in the text with its decimal equivalent
		text = strings.ReplaceAll(text, match[0], strconv.FormatInt(decimal, 10))
	}
	// Return the modified text with binary numbers replaced by decimal equivalents
	return text
}

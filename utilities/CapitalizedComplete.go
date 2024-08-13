package utilities

import (
	"strconv" 
	"strings" 
)

func CapitalizePrevious(s string) string {
	// Split the input string into individual words
	words := strings.Split(s, " ")

	//loop through the words slice to find "(cap" and "(cap,"
	for i := 0; i < len(words); i++ {
		// Check if the current word contains the characters "(cap" and "(cap,"
		if strings.Contains(words[i], "(cap") {
			
			if strings.Contains(words[i], "(cap,") {
				// Convert the next word (the one after "(cap,") into a number
				num, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])

				
				for j := i - num; j < i; j++ {
					
					words[j] = strings.ToUpper(string(words[j][0]))+words[j][1:]
				}

				// remove word cap and number from the sentence
				words = append(words[:i], words[i+2:]...)
			} else {
				//if the word doesn't contain "(cap," make the previous word Capitalized
				//and remove the current word from the sentence
				words[i-1] = strings.ToUpper(string(words[i-1][0]))+words[i-1][1:]

				
				words = append(words[:i], words[i+1:]...)
			}
		}
	}

	// Join the words back together into a single string with spaces between each word
	return strings.Join(words, " ")
}

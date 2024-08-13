package utilities

import (
	"strconv" 
	"strings" 
)


func LowerCase(s string) string {
	// Split the input string into individual words
	words := strings.Split(s, " ")

	//loop through the words slice to find "(low" and "(low,"
	for i := 0; i < len(words); i++ {
		
		if strings.Contains(words[i], "(low") {
			
			if strings.Contains(words[i], "(low,") {
				// Convert the next word (the one after "(low,") into a number
				num, _ := strconv.Atoi(words[i+1][:len(words[i+1])-1])

				//create another loop to iterate over the words to make lowercase
				for j := i - num; j < i; j++ {
					
					words[j] = strings.ToLower(words[j])
				}

				// remove word low and number from the sentence
				words = append(words[:i], words[i+2:]...)
			} else {
				//if the word doesn't contain "(low," make the previous word lowercase
				//and remove the current word from the sentence
				words[i-1] = strings.ToLower(words[i-1])

			
				words = append(words[:i], words[i+1:]...)
			}
		}
	}

	// Join the words back together into a single string with spaces between each word
	return strings.Join(words, " ")
}

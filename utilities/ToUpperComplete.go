package utilities

import (
	"strings"
	"strconv"
)

func UppercaseComplete(s string) string{
	//split the input string into a slice of words
	words:=strings.Split(s, " ")

	//loop through the words slice to find "(up" and "(up,"
	for i:=0; i<len(words);i++{
		if strings.Contains(words[i], "(up"){
			if strings.Contains(words[i], "(up,"){
				//convert the next word (the one after"(up,") into a number
				number,_:=strconv.Atoi(words[i+1][:len(words[i+1])-1])
			//create another loop to iterate over the words to make uppercase
			for k:=i-number;k<i;k++{
				words[k]=strings.ToUpper(words[k])

			}
			// remove word up and number from the sentence
			words=append(words[:i], words[i+2:]...)
			
			}else{
				//if the word doesn't contain "(up," make the previous word uppercase
				//and remove the current word from the sentence
				words[i-1]=strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
			}
		}
	}
	//join the words back together into single sentence

	return strings.Join(words, " ")

}
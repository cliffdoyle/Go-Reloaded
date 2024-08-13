package utilities

import "strings"

func AtoAn(sentence string) string {
	//split the sentece into individual words then iterate to find vowels
	words := strings.Fields(sentence)
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a" || words[i] == "A") && strings.ContainsAny("aeiouhAEIOUH", string(words[i+1][0])) {
			if words[i] == "a" {
				words[i] = "an"

			}else{
				words[i]="An"
			}

		}
	}
	return strings.Join(words, " ")
}

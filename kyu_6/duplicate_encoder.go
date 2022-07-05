// Link to description: https://www.codewars.com/kata/54b42f9314d9229fd6000d9c/train/go
package main

import (
	"fmt"
	"strings"
)

func duplicateEncode(word string) string {
	word = strings.ToLower(word)
	var result string
	for posA, charA := range word {
		rep := false
		for posB, charB := range word {
			if posA != posB && charA == charB {
				rep = true
				break
			}
		}
		if rep == false {
			result += "("
		} else {
			result += ")"
		}
	}

	return result
}

func duplicateEncodeMapApproach(word string) string {
	var result string
	word = strings.ToLower(word)
	var encoder = make(map[uint8]string)

	for i := 0; i < len(word); i++ {
		_, exists := encoder[word[i]]
		if exists {
			encoder[word[i]] = ")"
		} else {
			encoder[word[i]] = "("
		}
	}

	for i := 0; i < len(word); i++ {
		result += encoder[word[i]]
	}

	return result
}

func main() {
	fmt.Println(duplicateEncode("Hello()"))
	fmt.Println(duplicateEncodeMapApproach("Hello()"))
}

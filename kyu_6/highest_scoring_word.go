// Link to description: https://www.codewars.com/kata/57eb8fcdf670e99d9b000272/train/go
package main

import (
	"fmt"
	"strings"
)

func high(s string) string {
	max := 0
	var result string
	wordArr := strings.Split(s, " ")
	for _, word := range wordArr {
		current := 0
		for _, char := range word {
			current += int(char) - 96
		}
		if current > max {
			max = current
			result = word
		}
	}
	return result
}

func main() {
	fmt.Println(high("man i need a taxi up to ubud"))
}

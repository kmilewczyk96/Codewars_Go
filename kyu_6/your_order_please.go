// Link to description: https://www.codewars.com/kata/55c45be3b2079eccff00010f/train/go
package main

import (
	"fmt"
	"strings"
)

func order(sentence string) string {
	split := strings.Split(sentence, " ")
	ordered := make([]string, len(split))

	for i := 0; i < len(split); i++ {
		word := split[i]
		for j := 0; j < len(word); j++ {
			if word[j] >= 49 && word[j] <= 57 {
				ordered[word[j]-49] = word
				break
			}
		}
	}

	return strings.Join(ordered, " ")
}

func main() {
	fmt.Println(order("is2 Thi1s T4est 3a"))
}

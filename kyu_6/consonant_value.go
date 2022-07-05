// Link to description: https://www.codewars.com/kata/59c633e7dcc4053512000073/train/go
package main

import (
	"fmt"
	"strings"
)

func splitter(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func solve(str string) int {
	splitStr := strings.FieldsFunc(str, splitter)
	max := 0
	for i := 0; i < len(splitStr); i++ {
		current := 0
		piece := splitStr[i]
		for j := 0; j < len(piece); j++ {
			current += int(piece[j]) - 96
		}
		if current > max {
			max = current
		}
	}
	return max
}

func main() {
	fmt.Println(solve("zea"))
}

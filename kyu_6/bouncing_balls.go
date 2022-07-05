// Link to description: https://www.codewars.com/kata/5544c7a5cb454edb3c000047/train/go
package main

import "fmt"

func bouncingBall(h, bounce, window float64) int {
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}
	reps := 1 // Ball always passes the window on its way down.
	for true {
		h *= bounce
		if h > window {
			reps += 2
		} else {
			break
		}
	}
	return reps
}

func main() {
	fmt.Println(bouncingBall(3, 0.66, 1.5))
}

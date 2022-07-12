package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IPsBetween(start, end string) int {
	var addressesCount int64
	startArr := strings.Split(start, ".")
	endArr := strings.Split(end, ".")

	var multiplier int64 = 1
	for i := 3; i >= 0; i-- {
		minuend, _ := strconv.ParseInt(endArr[i], 10, 64)
		subtrahend, _ := strconv.ParseInt(startArr[i], 10, 64)
		addressesCount += (minuend - subtrahend) * multiplier
		multiplier *= 256
	}

	return int(addressesCount)
}

func main() {
	fmt.Println(IPsBetween("10.255.0.5", "15.0.0.1"))
}

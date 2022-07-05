// Link to description: https://www.codewars.com/kata/554a44516729e4d80b000012/train/go
package main

import "fmt"

func nbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	if startPriceOld >= startPriceNew {
		return [2]int{0, startPriceOld - startPriceNew}
	}
	month := 0
	var savings float64
	currentPriceOld := float64(startPriceOld)
	currentPriceNew := float64(startPriceNew)

	for currentPriceNew > (currentPriceOld + savings) {
		month += 1
		if month%2 == 0 {
			percentLossByMonth += 0.5
		}
		currentPriceOld *= (100 - percentLossByMonth) / 100
		currentPriceNew *= (100 - percentLossByMonth) / 100
		savings += float64(savingperMonth)
	}

	moneyLeft := int(currentPriceOld - currentPriceNew + savings + 0.5) // 0.5 added to force rounding instead of trunc.

	return [2]int{month, moneyLeft}
}

func main() {
	fmt.Println(nbMonths(8000, 12000, 500, 1.0))
}

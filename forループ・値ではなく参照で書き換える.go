package main

import "fmt"

func test() {
	var intQuantity int
	fmt.Scan(&intQuantity)
	intArray := make([]int, intQuantity)

	for i := 0; i < intQuantity; i++ {
		fmt.Scan(&intArray[i])
	}

	min := -1
	divisionCount := 0
	for i := 0; i < intQuantity; i++ {
		for intArray[i]%2 == 0 {
			intArray[i] /= 2
			divisionCount++
		}
		if min < 0 || min > divisionCount {
			min = divisionCount
		}
		divisionCount = 0

	}

	fmt.Printf("%v", min)
}

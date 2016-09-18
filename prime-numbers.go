package main

import (
	"fmt"
	"math"
)

func main() {
	firstValue := 2

	primeNumbers := make([]int, 1, 100)
	primeNumbers[0] = firstValue

	currentValue := firstValue

	for len(primeNumbers) < 100 {
		currentValue++

		isDivisible := isDivisible(currentValue, primeNumbers)

		if !isDivisible {
			primeNumbers = append(primeNumbers, currentValue)
			fmt.Println(currentValue)
		}
	}

	fmt.Println(primeNumbers)
}

func isDivisible(number int, primeNumbers []int) bool {
	for _, value := range primeNumbers {
		if math.Mod(float64(number), float64(value)) == 0 {
			return true
		}
	}
	return false
}

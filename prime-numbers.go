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

		isDivisible := false
		for _, value := range primeNumbers {
			if math.Mod(float64(currentValue), float64(value)) == 0 {
				isDivisible = true
				break
			}
		}

		if !isDivisible {
			primeNumbers = append(primeNumbers, currentValue)
			fmt.Println(currentValue)
		}
	}

	fmt.Println(primeNumbers)
}

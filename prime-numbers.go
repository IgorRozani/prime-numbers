package main

import (
	"fmt"
	"math"
)

func main() {
	const firstPrimeNumber = 2

	primeNumbers := make([]int, 1, 100)
	primeNumbers[0] = firstPrimeNumber

	currentValue := firstPrimeNumber

	fmt.Println("Calculating first 100 prime numbers...")

	for len(primeNumbers) < 100 {
		currentValue++

		isDivisible := isDivisible(currentValue, primeNumbers)

		if !isDivisible {
			primeNumbers = append(primeNumbers, currentValue)
		}
	}

	fmt.Println("The prime numbers are:", primeNumbers)
}

func isDivisible(number int, primeNumbers []int) bool {
	for _, value := range primeNumbers {
		if math.Mod(float64(number), float64(value)) == 0 {
			return true
		}
	}
	return false
}

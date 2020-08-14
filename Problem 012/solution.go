package main

import (
	"fmt"
	"math"
	// "time"
)

func sqrt(num int) float64 {
	res := math.Sqrt(float64(num))

	return res
}

func divisors(num int) int {
	var numOfDivisor = 0
	for i := 1; i <= int(math.Ceil(sqrt(num))); i++ {
		if num % i == 0 {
			numOfDivisor += 2
		}

		if i * i == num {
		numOfDivisor -= 1
		}
	}

	return numOfDivisor
}

func main() {
	// start := time.Now()
	var nthTriangle, cnt int

	for num := 1; ;num++ {
		nthTriangle = (num * (num + 1)) / 2

		if num % 2 == 0 {
			cnt = divisors(num / 2) * divisors(num + 1)
		} else {
			cnt = divisors(num) * divisors((num + 1) / 2)
		}

		if cnt >= 500 {
			fmt.Println(nthTriangle)
			break
		}
	}

	// elapsed := time.Since(start)
	// fmt.Printf("\nSolution took: %s\n", elapsed)
}

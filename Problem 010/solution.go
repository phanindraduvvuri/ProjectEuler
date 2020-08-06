package main

import (
	"fmt"
	"math"
)

func sieveOfErathosthenes() []int {
	limit := 1000000
	sieve := make([]bool, limit+1)

	for i := 0; i <= limit; i++ {
		if (i%2 == 0) && (i != 2) {
			sieve[i] = true
		} else {
			sieve[i] = false
		}
	}

	for i := 3; i <= int(math.Sqrt(float64(limit))); i += 2 {
		if !sieve[i] {
			for j := i * i; j < limit; j++ {
				sieve[j] = true
			}
		}
	}

	sumOfPrimes := []int{0, 1}
	sumSofar := 0
	for i := 2; i < limit+1; i++ {
		if !sieve[i] {
			sumSofar += i
		}
		sumOfPrimes = append(sumOfPrimes, sumSofar)
	}

	return sumOfPrimes
}

func main() {
	var cases int
	sumOfPrimes := sieveOfErathosthenes()
	fmt.Scan(&cases)
	for c := 0; c < cases; c++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(sumOfPrimes[n])
	}
}

package main

import (
	"fmt"
)


func multiples_sum(num, k int) int {
	common_diff := num / k
	return k * ((common_diff) * (common_diff + 1)) / 2
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var N int
		fmt.Scan(&N)
		result := multiples_sum(N - 1, 3) + multiples_sum(N - 1, 5) - multiples_sum(N - 1, 15)
		fmt.Println(result)
	}
}

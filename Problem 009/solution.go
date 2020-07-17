package main

import (
	"fmt"
)

func  get_triplet_product(n int) int {
	prod := -1
	for a := 3; a < n / 3; a++ {
		b := ((n * n) - (2 * a * n)) / ((2 * n) - (2 * a))
		c := n - a - b

		if ((c * c) == (a * a) + (b * b)) {
			temp := a * b * c
			if temp > prod {
				prod = temp
			}
		
		}
	}

	return prod
}

func main()  {
	var cases int
	fmt.Scan(&cases)
	for c := 0; c < cases; c++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(get_triplet_product(n))
		
	}
}

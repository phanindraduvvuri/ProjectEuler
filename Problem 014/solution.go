package main

import "fmt"

func collatzSequenceUtil(n int, collLenMap map[int]int) int {
	if _, ok := collLenMap[n]; ok {
		return collLenMap[n]
	}

	if n == 1 {
		collLenMap[n] = 1
	} else if n%2 == 0 {
		collLenMap[n] = 1 + collatzSequenceUtil(n/2, collLenMap)
	} else {
		collLenMap[n] = 1 + collatzSequenceUtil((3*n)+1, collLenMap)
	}

	return collLenMap[n]
}

func collatzSequence(n int) (int, int) {
	lenght := -1
	num := 0
	collLenMap := make(map[int]int)

	collatzSequenceUtil(n, collLenMap)

	for i := 1; i < n; i++ {
		if _, ok := collLenMap[i]; !ok {
			collatzSequenceUtil(i, collLenMap)
		}

		currentLen := collLenMap[i]

		if lenght < currentLen {
			lenght = currentLen
			num = i
		}
	}

	return num, lenght
}

func main() {
	fmt.Println(collatzSequence(1000000))
}

package main

import "fmt"

func isPalindrom(num int) bool {
    reversed := 0
    number := num

    for num > 0 {
        reversed = reversed * 10 + num % 10
        num /= 10
    }

    return reversed == number
}

func main() {
    var t int
    fmt.Scan(&t)
    for c := 0; c < t; c++ {
        var n, prod int
        fmt.Scan(&n)

        largestPrime := -1
        for i := 990; i > 99; i-=11 {
            for j := 999; j > 99; j-- {
                prod = i * j
                if prod > largestPrime && isPalindrom(prod) && prod < n {
                    largestPrime = prod
                    break
                } else if prod < largestPrime {
                    break
                }
            }
        }
        fmt.Println(largestPrime)
    }
}

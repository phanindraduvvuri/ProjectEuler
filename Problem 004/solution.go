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
    for i := 0; i < t; i++ {
        var n, prod int
        fmt.Scan(&n)
        largest_prime := -1

        for j := 990; j <= 100; j-=11 {
            for k := 999; k <= 100; k-- {
                prod = j * k
                if prod > largest_prime && isPalindrom(prod) && prod < n {
                    largest_prime = prod
                    break
                } else if prod < largest_prime {
                    break
                }
            }
        }
        fmt.Println(largest_prime)
    }
}

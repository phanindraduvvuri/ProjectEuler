package main

import (
    "fmt"
)

func sumOfNumbers(n int) int {
    return (n * (n + 1)) / 2
}

func sumOfSquares(n int) int {
    return (n * (n + 1) * (2 * n + 1)) / 6
}

func pow(base, exp int) int {
    var result = 1
    for i := 0; i < exp; i++ {
        result *= base
    }

    return result
}

func main()  {
    var tests int
    fmt.Scan(&tests)
    for i := 0; i < tests; i++ {
        var n int
        fmt.Scan(&n)

        result := pow(sumOfNumbers(n), 2) - sumOfSquares(n)
        fmt.Println(result)
    }
}

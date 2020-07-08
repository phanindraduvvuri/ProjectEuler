package main

import "fmt"

func max(val1, val2 int) int {
    if val1 > val2 {
        return  val1
    } else {
        return val2
    }
}

func prod(number string) int {
    prod := 1
    for _, digit := range number {
        prod *= int(digit - '0')
    }

    return prod
}

func main() {
    var cases int
    fmt.Scan(&cases)

    for c := 0; c < cases; c++ {
        var n, k, max_prod int
        fmt.Scan(&n, &k)

        var number string
        fmt.Scan(&number)

        max_prod = -1
        for i := 0; i <= n - k; i++ {
            val := prod(number[i:i + k])
            max_prod = max(val, max_prod)
        }

        fmt.Println(max_prod)
    }
}

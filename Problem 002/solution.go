package main

import "fmt"

func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        a := 1
        b := 2

        total_sum := 0
        for a < n {
            if a % 2 == 0 {
                total_sum += a
            }
            a, b = b, a + b
        }
        fmt.Println(total_sum)
    }
}

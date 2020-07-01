package main

import (
    "fmt"
    "math"
)

func main()  {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {

        var n, largestPrimeFactor int
        fmt.Scan(&n)

        largestPrimeFactor = 2
        for n % 2 == 0 {
            n = n / 2
        }

        for i := 3; i <= int(math.Sqrt(float64(n))); i++ {
            for n % i == 0 {
                if i > largestPrimeFactor {
                    largestPrimeFactor = i
                }
                n /= i
            }
        }

        if n > 2 {
            largestPrimeFactor = n
        }

        fmt.Println("Largest Prime Factor: ", largestPrimeFactor)
    }
}

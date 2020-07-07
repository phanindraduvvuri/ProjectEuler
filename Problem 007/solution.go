package main

import "fmt"

const (
    MAX_SIZE = 105000
)

var (
    Primes = make([]int, 0)
)

func sieveOfErastosthenes() {
    var isPrime = make([]bool, MAX_SIZE)

    for i := range isPrime {
        isPrime[i] = true
    }

    for p := 2; p * p < MAX_SIZE; p++ {
        if isPrime[p] {
            for i := p * p; i < MAX_SIZE; i += p {
                isPrime[i] = false
            }
        }
    }

    for p := 2; p < MAX_SIZE; p++ {
        if isPrime[p] {
            Primes = append(Primes, p)
        }
    }
}

func nthPrime(n int) int {
    return Primes[n - 1]
}

func main() {
    var cases int
    fmt.Scan(&cases)
    sieveOfErastosthenes()

    for c := 0; c < cases; c++ {
        var n int
        fmt.Scan(&n)
        fmt.Println(nthPrime(n))
    }
}

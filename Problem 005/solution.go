package main

import "fmt"

func gcd(a, b int) int {
    if b == 0 {
        return a
    } else {
        return gcd(b, a % b)
    }
}

func lcm(sli []int) int {
    result := 1
    for _, ele := range sli {
        result = result * ele / gcd(result, ele)
    }
    return result
}


func main()  {
    var tests int
    fmt.Scan(&tests)
    for c := 0; c < tests; c++ {
        var n int
        sli := []int{}
        fmt.Scan(&n)
        for i := 1; i <= n; i++ {
            sli = append(sli, i)
        }

        fmt.Println(lcm(sli))
    }

}

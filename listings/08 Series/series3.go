package main

import "fmt"

func main() {
    var number float32
    var sum float32 = 0
    for i := 0; i < 10; i++ {
        fmt.Scan(&number)
        sum += number
    }
    aMean := sum / 10
    fmt.Printf("aMean = %.2f\n", aMean)
}
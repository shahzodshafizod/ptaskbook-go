package main

import "fmt"

func main() {
    var a, b int
    fmt.Print("A = ")
    fmt.Scan(&a)
    fmt.Print("B = ")
    fmt.Scan(&b)
    result := (a + b) % 2 != 0
    fmt.Printf("result = %t\n", result)
}
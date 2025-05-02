package main

import "fmt"

func main() {
    a, b := 10, 5

    fmt.Println("Addition:", a+b)
    fmt.Println("Multiplication:", a*b)
    fmt.Println("Comparison:", a > b)
    fmt.Println("Logical AND:", (a > 0) && (b > 0))
}
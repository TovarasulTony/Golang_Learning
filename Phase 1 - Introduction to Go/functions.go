package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

//multiple returns
func swap(a, b string) (string, string) {
    return b, a
}

func main () {
    sum := add(5, 7)
    fmt.Println("Sum:", sum)

    x,y := swap("Hello", "World")
    fmt.Println(x,y)
}
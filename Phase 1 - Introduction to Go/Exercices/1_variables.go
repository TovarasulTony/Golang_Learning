package main

import "fmt"

/* Ex.
Declare and initialize variables of different types (int, float64, string, bool).
Declare a constant and try to modify it (observe the error).
Print all values.
*/

func main() {
    var a int = 100
    var b float64 = 3.4
    var c string = "Toss a coin foru your..."
    var d bool = false

    const e int = 11
    // e = 20 ERROR: cannot assign to e (neither addressable nor a map index expression)
    fmt.Println("a = ", a)
    fmt.Println("b = ", b)
    fmt.Println("c = ", c)
    fmt.Println("d = ", d)
    fmt.Println("e = ", e)
}
package main

import "fmt"

/* Ex.
If a number is even, print "Even"; otherwise, print "Odd".
*/

func main() {
    var num = 11
    if num % 2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}
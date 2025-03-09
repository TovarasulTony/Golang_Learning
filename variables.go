package main

import "fmt"

func main() {
    // 1. Using var (Explicit type)
    var name string = "Golang"

    // 2. Type inference (Go determines the type automatically)
    var version = 1.20

    // 3. Short declaration (inside functions only)
    language := "Go"

    fmt.Println("Name", name)
    fmt.Println("Version", version)
    fmt.Println("Language", language)

    // Multiple Variable Declaration
    var a, b, c int = 1, 2, 3
    x, y, z := "Hello", true, 3.14
    fmt.Println("a", a)
    fmt.Println("b", b)
    fmt.Println("c", c)
    fmt.Println("z", z)
    fmt.Println("y", y)
    fmt.Println("x", x)

    // Constants
    const Pi = 3.14159
    const Greeting = "Hello, Go!"
    
    fmt.Println(Greeting)
    fmt.Println("Value of Pi:", Pi)
}
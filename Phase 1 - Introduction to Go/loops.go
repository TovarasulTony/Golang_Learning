package main

import "fmt"

func forLoop() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Iteration:", i)
    }
}

// For Loop Without Initialization
// No while in GO, this variance can be used liek while
func forVariance() {
    i := 1
    for i <= 5 {
        fmt.Println(i)
        i++
    }
}

func infiniteFor() {
    // can use break to end
    for {
        fmt.Println("Running forever...")
    }
}

func forRange() {
    nums := []int{10, 20, 30, 40}

    for index, value := range nums {
        fmt.Println("Index:", index, "Value:", value)
    }
}

// ignoring index
func forRange2() {
    nums := []int{10, 20, 30, 40}

    for _, value := range nums {
        fmt.Println(value)
    }
}


func main() {
    forLoop()
    forVariance()
    //infiniteFor()
    forRange()
    forRange2()
}
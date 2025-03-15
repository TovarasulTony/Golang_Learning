package main

import "fmt"

func typeConversion() {
    var a int = 10
    var b float64 = 5.5
    sum := float64(a) + b // Type conversion required

    fmt.Println("Integer:", a)
    fmt.Println("Float:", b)
    fmt.Println("Sum:", sum)
}

func stringExample() {
    var message string = "Hello, Go!"
    fmt.Println(message)

    // String concatenation
    name := "Alice"
    greeting := "Hello, " + name
    fmt.Println(greeting)

    // Get string length
    fmt.Println("Length:", len(message))

    // Access characters (returns byte, not rune)
    fmt.Println("First character:", string(message[0]))
}

func arrayExample() {
    fmt.Println("\n---arrayExample---")

    // short notation
    // arr := [...]int{10, 20, 30}
    var arr [3]int = [3]int{10, 20, 30} // Fixed size 3
    fmt.Println("Array:", arr)

    arr[1] = 50 // Modify an element
    fmt.Println("Updated Array:", arr)

    // Length of array
    fmt.Println("Length:", len(arr))
}

// Slices are dynamic-sized arrays in Go and are more commonly used than arrays.
func slicesExample() {
    fmt.Println("\n---slicesExample---")

    nums := []int{1, 2, 3, 4, 5} // Slice (dynamic)
    fmt.Println("Slice:", nums)

    nums = append(nums, 6) // Append new element
    fmt.Println("After Append:", nums)

    part := nums[1:4] // Slice a portion
    fmt.Println("Sliced Portion:", part)

    fmt.Println("Length:", len(nums), "Capacity:", cap(nums))
}

func mapExample() {
    /*
    delete

    */
    fmt.Println("\n---mapExample---")
    var person map[string]int = map[string]int {
        "Alice": 25,
        "Bob":   30,
    }
    fmt.Println("Map:", person)

    person["Charlie"] = 35 // Add new key-value pair
    fmt.Println("Updated Map:", person)

    // Access value
    age := person["Alice"]
    fmt.Println("Alice's Age:", age)

    // Delete a key
    delete(person, "Bob")
    fmt.Println("After Deletion:", person)
}

func main() {
    typeConversion()
    stringExample()
    arrayExample()
    slicesExample()
    mapExample()
}
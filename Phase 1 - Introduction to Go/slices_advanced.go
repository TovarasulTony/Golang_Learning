package main

import "fmt"

/* Slices are dynamic-sized arrays in Go and are more commonly used than arrays.
   A slice is a view over an array
   Internally, a slice in Go does not store the actual data. Instead, it stores:
   - A pointer to the underlying array
   - The slice length
   - The slice capacity
*/
func slicesAppend() {
    fmt.Println("\n---slicesAppend---")
    slice1 := []int{1, 2, 3}
    slice2 := []int{4, 5, 6}
    slice1 = append(slice1, slice2...) // Appends all elements from slice2
    fmt.Println(slice1) // Output: [1 2 3 4 5 6]
    fmt.Println("Length:", len(slice1), "Capacity:", cap(slice1))
}

func viewOverArray() {
    fmt.Println("\n---viewOverArray---")
    arr := [5]int{10, 20, 30, 40, 50} // Array
    slice := arr[1:4] // Slice referencing elements 20, 30, 40
    fmt.Println("Slice:", slice) // Output: [20 30 40]
    slice[0] = 100 // Modifying slice affects the original array
    fmt.Println("\nModified Slice:", slice) // Output: [100 30 40]
    fmt.Println("Original Array:", arr) // Output: [10 100 30 40 50]

    slice = append(slice, 60)
    slice[1] = 100
    fmt.Println("\nModified Slice:", slice)
    fmt.Println("Original Array:", arr)

    slice = append(slice, 70)
    slice[2] = 100
    fmt.Println("\nModified Slice:", slice)
    fmt.Println("Original Array:", arr)
}

func sliceExample() {
    fmt.Println("\n---sliceExample---")
    arr := [3]int{1, 2, 3}
    slice := arr[:] // Slice covering the entire array
    slice = append(slice, 4) // Exceeds array capacity, creates new array

    fmt.Println(slice) // Output: [1 2 3 4]
    fmt.Println(arr)   // Output: [1 2 3] (unchanged!)
}

func sliceCopy() {
    fmt.Println("\n---sliceCopy---")
    arr := [3]int{1, 2, 3}
    slice := make([]int, len(arr)) // Create new slice with same length
    copy(slice, arr[:]) // Copy data

    slice[0] = 100 // Modify slice
    fmt.Println("Original Array:", arr) // Output: [1 2 3] (unchanged!)
    fmt.Println("New Slice:", slice)    // Output: [100 2 3]
}

func main() {
    slicesAppend()
    viewOverArray()
    sliceExample()
    sliceCopy()
}
package main

import "fmt"

func genericIf() {
    age := 18
    if age >= 18 {
        fmt.Println("You are an adult.")    
    } else {
        fmt.Println("You are a minor.")
    }
}

func shortIfDeclaration() {
    // num is scoped
    if num:= 10; num%2 == 0 {
        fmt.Println("Even Number")
    } else {
        fmt.Println("Odd Number")
    }
}

func switchCase() {
    day := "Tuesday"

    switch day {
    case "Monday":
        fmt.Println("Start of the week")
    case "Friday":
        fmt.Println("Weekend is near!")
    default:
        fmt.Println("It's a regular day.")
    }
}

func switchFallthrough() {
    /* Important Notes on fallthrough
    1. fallthrough only moves to the next case, even if the condition does not match.
    2. It cannot be used in the default case (it will cause a compile-time error).
    3. It does not check the next case condition—it just continues execution.
    */
    num := 2

    switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
        fallthrough // Forces execution of the next case
    case 3:
        fmt.Println("Three") // This runs due to fallthrough
    default:
        fmt.Println("Other number")
    }
}

func switchCaseLists() {
    day := "Saturday"

    switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("It's a weekday")
    case "Saturday", "Sunday":
        fmt.Println("It's the weekend!")
    default:
        fmt.Println("Not a valid day")
    }
}

func switchCaseListsBooleanExpression() {
    score := 85

    switch {
    case score >= 90:
        fmt.Println("Grade: A")
    case score >= 80 && score < 90:
        fmt.Println("Grade: B")
    case score >= 70 && score < 80:
        fmt.Println("Grade: C")
    default:
        fmt.Println("Grade: F")
    }
}

func main() {
    genericIf()
    shortIfDeclaration()

    switchCase()
    switchFallthrough()
    switchCaseLists()
    switchCaseListsBooleanExpression()
}
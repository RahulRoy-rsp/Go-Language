package main

import "fmt"

func main() {
    
    // if statement: to check if a number is less than 5
    x := 4
    if x < 5 {
        fmt.Println("YES, num is less than 5")          // Output: YES, num is less than 5
    }
    // it will print:
    //              YES, num is less than 5


    // if-else statement: to check if a number is even or odd
    n := 12
    if (n%2 == 0) {
        fmt.Println("num is even number")          // Output: num is even number
    } else {
        fmt.Println("num is odd number")            // Output: num is even number
    }
    // it will print:
    //              num is even number


    // if - else if - else statement: to check if a number is positive, zero or negative.
    num := 10
    if num > 0 {
        fmt.Println("num is positive") // Output: num is positive
    } else if num == 0 {
        fmt.Println("num is zero") // Output: num is zero
    } else {
        fmt.Println("num is negative") // Output: num is negative
    }
    // it will print: 
    //              num is positive


    // switch statement:
    num = 74
    switch num {
    case 0:
        fmt.Println("num is zero") // Output: num is zero
    case 1, 2, 3, 4, 5:
        fmt.Println("num is between 1 to 5") // Output: num is between 6 to 10
    case 6, 7, 8, 9, 10:
        fmt.Println("num is between 6 to 10") // Output: num is between 6 to 10
    default:
        fmt.Println("num is out of the range") // Output: num is out of the range
    }
    // it will print: 
    //              num is out of the range


    // Nested if statement: to check if gender is male or female and then check their age
    gender := "F"
    age := 21
    if gender == "F" {
        if age >= 21 {
            fmt.Println("You are eligible for marriage.")
        } else {
            fmt.Println("You are not eligible for marriage.")
        }
    } else if gender == "M" {
        if age >= 22 {
            fmt.Println("You are eligible for marriage.")
        } else {
            fmt.Println("You are not eligible for marriage.")
        }
    }   
    // it will print:
    //              You are eligible for marriage.
}

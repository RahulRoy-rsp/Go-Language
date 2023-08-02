package main

import "fmt"

func main() {
    // Arithmetic operators on a, b
    a, b := 10, 3
    fmt.Println("Addition:", a+b)                       // op: Addition: 13
    fmt.Println("Subtraction:", a-b)                    // op: Subtraction: 7
    fmt.Println("Multiplication:", a*b)                 // op: Multiplication: 30
    fmt.Println("Division:", a/b)                       // op: Division: 3
    fmt.Println("Remainder:", a%b)                      // op: Remainder: 1

    // Relational operators
    fmt.Println("is a Equal b:", a == b)                // op: is a Equal b: false
    fmt.Println("is a Not equal b:", a != b)            // op: is a Not equal b: true
    fmt.Println("is a Greater than b:", a > b)          // op: is a Greater than b: true
    fmt.Println("is a Less than b:", a < b)             // op: is a Less than b: false

    // Logical operators
    fmt.Println("Logical AND:", true && false)          // op: Logical AND: false
    fmt.Println("Logical OR:", true || false)           // op: Logical OR: true
    fmt.Println("Logical NOT:", !true)                  //  op: Logical NOT: false
    
    // Bitwise operators
    fmt.Println("Bitwise AND:", a&b)                    // op: Bitwise AND: 2
    fmt.Println("Bitwise OR:", a|b)                     // op: Bitwise OR: 11
    fmt.Println("Bitwise XOR:", a^b)                    // op: Bitwise XOR: 9

    fmt.Println("Bitwise left shift:", a<<2)            // op: Bitwise left shift: 40
    fmt.Println("Bitwise right shift:", a>>2)           // op: Bitwise right shift: 2


    // Assignment operators
    x := 5
    x += 2
    fmt.Println("x += 2:", x)                           // op: x += 2: 7
    x -= 2
    fmt.Println("x -= 2:", x)                           // op: x -= 2: 5
    x *= 2
    fmt.Println("x *= 2:", x)                           // op: x *= 2: 10
    x /= 2
    fmt.Println("x /= 2:", x)                           // op: x /= 2: 5


}

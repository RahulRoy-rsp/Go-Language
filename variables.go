package main
import "fmt"

func main() {

    // \n is used to break the current line and go to the next line
    
    // Implicit declaration

        // Numeric example
            // integer
    var roll1 int
    fmt.Printf("roll1 is type of %T\n", roll1)
    
            // float
    var marks float32
    fmt.Printf("marks is type of %T\n", marks)

        // Boolean example
    var isPassed bool
    fmt.Printf("isPassed is type of %T\n", isPassed)
    
        // String example
    var name string
    fmt.Printf("name is type of %T\n", name)

    // Explicit declaration
    var roll = 45
    var grade = 7.4
    var isFailed = false
    var Name = "Alex"
    
    fmt.Printf("roll is type of %T\n", roll)
    fmt.Printf("grade is type of %T\n", grade)
    fmt.Printf("isFailed is type of %T\n", isFailed)
    fmt.Printf("Name is type of %T\n", Name)


    // Dynamic declaration
    newRoll := 12
    newGrade := 44.7
    married := false
    newName := "Malmo"
    
    fmt.Printf("newRoll is type of %T\n", newRoll)
    fmt.Printf("newGrade is type of %T\n", newGrade)
    fmt.Printf("married is type of %T\n", married)
    fmt.Printf("newName is type of %T\n", newName)

    // Multiple variable declaration in one line
    var p, q, r = 7, 8, "nine"
    x, y, z := 2, "three", 4 
    
    fmt.Println("p is:", p)
    fmt.Println("q is:", q)
    fmt.Println("r is:", r)    
    fmt.Println("x is:", x)
    fmt.Println("y is:", y)
    fmt.Println("z is:", z)
}

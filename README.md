# GO Language
Concepts about the GoLang I learned so far...
---
#### NOTE: GO is a Case Sensitive Language.
---
- **package main**
  - mandatory line.
  - It is the starting point to the program.
  - Go always runs with the package.

- **import "fmt"**
  - preprocessor command, meaning it commands Go compiler to include all the files from `fmt` package.
  - To import a package we usually write: `import 'package-name'`.

- **funct main()**
  - main function, where the compiler `starts the execution`.

- **Println()**
  - it is a method available in `fmt` package to `print a line`.
 
## [Click to see a Go Program to print a line](https://github.com/RahulRoy-rsp/Go-Language/blob/main/printing.go)
---
### Data Types Available:
    1. Numeric: 
       - These are numeric data which can be any number value.
       - int contains number without decimals. (Eg: 74)
       - float contains number with decimals. (Eg: 74.7)
    
    2. String
       - A string type represents the set of string values. (Eg: 'roy')
       - It means that we use alpha-numeric characters.
       - They are immutable, which means we can't change the contents of a string once its created. 
    
    3. Boolean
        - They contain only two values.
        - Either `true` or `false`. 
    
    4. Derived
       - These are user defined data types like Array, Structure, Pointers, etc.
---
## VARIABLES
### Syntax:
  1. **var variable_name variable_datatype**
      - Here, var is a default keyword to declare a variable followed by variable name and its datatype.
      - `var studentName string`
      - In above example the name of variable is `studentName` and its dataype is `string`.
  2. **var variable_name**
      - Here, var is a default keyword to declare a variable followed by variable name but without datatype because pf the explicit nature, the compiler will automatically understand the datatype of the variable based on its value.
      - `var studentName`
      - In above example the name of variable is `studentName` and its dataype will be `string`.
  3. **var variable_name = value**
      - Similarly like 2nd way, var is a default keyword to declare a variable followed by variable name and the value for that variable.
      - `var studentID = 5`
      - In above example the name of variable is `studentID` and its dataype will be automatically detected by compiler as an `integer`.
  4. **variable_name := value**
      - Lastly, to save time and declare a variable we write variable name followed by `:=` and then its value.
      - `studentMarks := 77`
      - In above example the name of variable is `studentMarks` and its dataype will be automatically detected by compiler as an `integer` with the value as `77`.

## [Click to see a Go Program for variable declaration](https://github.com/RahulRoy-rsp/Go-Language/blob/main/variables.go)
---

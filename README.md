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
## OPERATORS

In the Go programming language, operators are symbols or keywords used to perform various operations on variables, constants, and expressions. They are fundamental building blocks for performing arithmetic, logical, bitwise, and other types of operations in Go. Here are the different types of operators available in Go:

1. **Arithmetic Operators**:
   - Addition ( + ): to add the values.
     - Eg: a+b will give the addition of a and b.

   - Subtraction ( - ): to subtract the values.
     - Eg: a-b will give the value of a minus b.

   - Multiplication ( * ): to get the product of values.
     - Eg: a*b will give the value of a multiplied by b.

   - Division ( / ): to divide the values.
     - Eg: a/b will give the value of a divide by b.

   - Remainder (Modulus) ( % ): to get the remainder while dividing two values.
     - Eg: a%b will give the value of remainder while divinding a by b.

3. **Relational Operators**:
   - Equal to ( == ): checks if two values are equal or not, if equal then returns `TRUE` and if not then `False`.
     - Eg: a==b will check if the value of a is equals to b, if yes then returns TRUE and if not then FALSE.

   - Not equal to ( != ): checks if two values are not equal, if not equals then returns `TRUE` and if equals then `False`.
      - Eg: a!=b will check if the value of a is not equal to b, if yes then returns TRUE and if not then FALSE.

   - Greater than ( > ): checks if the value on left side is greater than the right side of the operand(>), if yes then returns `TRUE` and if not then `False`.
      - Eg: a>b will check if the value of a is greater than b, if yes then returns TRUE and if not then FALSE.

   - Greater than or equal to ( >= ): checks if the value on left side is greater or equal than the right side of the operand(>=), if yes then returns `TRUE` and if not then `False`.
      - Eg: a>=b will check if the value of a is greater or equal than b, if yes then returns TRUE and if not then FALSE.

   - Less than ( < ): checks if the value on left side is lesser than the right side of the operand(<), if yes then returns `TRUE` and if not then `False`.
      - Eg: a<b will check if the value of a is lesser than b, if yes then returns TRUE and if not then FALSE.

   - Less than or equal to ( <= ): checks if the value on left side is lesser or equal than the right side of the operand(<=), if yes then returns `TRUE` and if not then `False`.
      - Eg: a<=b will check if the value of a is lesser or equal than b, if yes then returns TRUE and if not then FALSE.
   
5. **Logical Operators**:
   - Logical AND ( && ): Performs AND operation on two operands, meaning that both the values on the side of the operator(&&) must be equal to return TRUE, otherwise it will return FALSE.
     -  Eg: Suppose `a` have value `TRUE` and `b` have value `FALSE`, then `a && b` will have value `FALSE`, because both the values are not equal.

   - Logical OR ( || ): Performs OR operation on two operands, meaning that if one of the values on the side of the operator(||) is TRUE, then it will return TRUE, otherwise it will return FALSE.
     -  Eg: Suppose `a` have value `TRUE` and `b` have value `FALSE`, then `a || b` will have value `TRUE`, because one of the value is TRUE.
    
   - Logical NOT ( ! ): Performs NOT operation, meaning it will negate/oppose the value.
     -  Eg: Suppose `a` have value `TRUE` and `b` have value `FALSE`, then `!a` will have value `False` and `!b` will have `TRUE`.
       
7. **Bitwise Operators**:
   - Bitwise AND ( & ): Performs bitwise AND operation on each corresponding bit of the operands.
     - Eg: 

   - Bitwise OR ( | ): Performs bitwise OR operation on each corresponding bit of the operands.
      - Eg:
   - Bitwise XOR ( ^ ): Performs bitwise XOR operation on each corresponding bit of the operands.
     - Eg:
       
   - Bitwise left shift ( << ): Shifts the bits of the left operand to the left by the number of positions specified by the right operand.
     - Eg:

   - Bitwise right shift ( >> ): Shifts the bits of the left operand to the right by the number of positions specified by the right operand.
     - Eg:

8. **Assignment Operators**:
   - Assignment ( = )
      - Eg:

   - Add and assign ( += )
      - Eg:

   - Subtract and assign ( -= )
      - Eg:

   - Multiply and assign ( *= )
      - Eg:

   - Divide and assign ( /= )
      - Eg:
        
*NOTE: There are many other operators available too, so you can refer them online On Google.*

## [Click to see a Go Program for some of Operators available](https://github.com/RahulRoy-rsp/Go-Language/blob/main/operators.go)
---
## Conditional Statements:

1. **if statement**:
  - Eg:
    
2. **if - else statement**:
  - Eg:
    
3. **if - else if - else statement**:
  - Eg:
    
4. **Switch statement**:
  - Eg:
     
5. **Nested if statement**:
  - Eg:
    
## [Click to see a Go Program for implementation of conditional statements](https://github.com/RahulRoy-rsp/Go-Language/blob/main/conditions.go)
---

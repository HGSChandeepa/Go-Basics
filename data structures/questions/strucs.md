## Problem Statement

You are tasked with creating a program that manages employee records using structs in Go. Your program should perform the following tasks:

1. Define a struct called `Employee` with the following fields:
   - `ID` (int)
   - `Name` (string)
   - `Position` (string)
   - `Salary` (float64)
2. Create a function called `printEmployeeDetails` that takes an `Employee` as an argument and prints out the employee's details in a formatted manner.
3. In the `main` function:
   - Create an instance of `Employee`.
   - Assign values to its fields.
   - Call the `printEmployeeDetails` function to display the employee's information.

## Requirements

- Use appropriate variable names and comments to explain your code.
- Ensure that your program compiles and runs without errors.

## Sample Output

```
Employee ID: 101
Name: John Doe
Position: Software Engineer
Salary: 75000.00
```

## Solution

```go
package main

import "fmt"

// Define the Employee struct
type Employee struct {
    ID       int
    Name     string
    Position string
    Salary   float64
}

// Function to print employee details
func printEmployeeDetails(emp Employee) {
    fmt.Printf("Employee ID: %d\n", emp.ID)
    fmt.Printf("Name: %s\n", emp.Name)
    fmt.Printf("Position: %s\n", emp.Position)
    fmt.Printf("Salary: %.2f\n", emp.Salary)
}

func main() {
    // Create an instance of Employee and assign values to its fields
    emp := Employee{
        ID:       101,
        Name:     "John Doe",
        Position: "Software Engineer",
        Salary:   75000.00,
    }

    // Call the function to print employee details
    printEmployeeDetails(emp)
}
```

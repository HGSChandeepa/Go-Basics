/*
Package main: This declares that the package name is main. In Go, the main package is special because it defines a standalone executable program.

Import "fmt": This imports the fmt package, which contains functions for formatted I/O, including printing to the console.

Func main() { ... }: This defines the main function, which is the entry point of the program. When the program is executed, the code inside this function is run.

Go run Hello.go: This is the command to run the Go program. It compiles the code and executes the resulting binary.
*/

package main

import "fmt"

func main() {
    // Print "Hello, World" multiple times
    fmt.Println("Hello, World") // Prints with a newline
    fmt.Println("Hello, World") // Prints with a newline
    fmt.Println("Hello, World") // Prints with a newline

    // Print "Hello, World" on the same line
    fmt.Print("Hello, World ")
    fmt.Print("Hello, World ")
    fmt.Print("Hello, World ")

    // ===============================================
    // Variables
    // ===============================================
    
    // Declaring variables with meaningful names
    var userName string = "John Doe"
    var userAge int = 25
    var userScore float32 = 25.5
    var userBalance float64 = 25.5
    var isUserFound bool = true

    // Print the type and value of each variable
    fmt.Printf("Type: %T Value: %v\n", userName, userName)
    fmt.Printf("Type: %T Value: %v\n", userAge, userAge)
    fmt.Printf("Type: %T Value: %v\n", userScore, userScore)
    fmt.Printf("Type: %T Value: %v\n", userBalance, userBalance)
    fmt.Printf("Type: %T Value: %v\n", isUserFound, isUserFound)

    // Declaring multiple variables of the same type
    var firstName, lastName string = "John", "Doe"
    fmt.Println(firstName, lastName)

    // Declaring multiple variables with different types
    var anotherName string = "Jane"
    var anotherAge int = 30
    fmt.Println(anotherName, anotherAge)

    // Short variable declaration
    newUserName := "Alice"
    fmt.Println(newUserName)
    fmt.Printf("Type: %T Value: %v\n", newUserName, newUserName)

}

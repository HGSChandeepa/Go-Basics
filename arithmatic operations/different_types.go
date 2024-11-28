package main

import "fmt"

func main() {
    // Declare integer and float variables
    var intNum int = 10
    var floatNum float64 = 3.5

    // Addition
    sum := float64(intNum) + floatNum // Convert int to float for addition
    fmt.Printf("Sum: %.2f\n", sum) // Output: Sum: 13.50

    // Subtraction
    difference := float64(intNum) - floatNum // Convert int to float for subtraction
    fmt.Printf("Difference: %.2f\n", difference) // Output: Difference: 6.50

    // Multiplication
    product := intNum * int(floatNum) // Convert float to int for multiplication
    fmt.Printf("Product: %d\n", product) // Output: Product: 10

    // Division
    quotient := float64(intNum) / floatNum // Convert int to float for division
    fmt.Printf("Quotient: %.2f\n", quotient) // Output: Quotient: 2.86

    // Modulo (only works with integers)
    var anotherIntNum int = 4
    remainder := intNum % anotherIntNum
    fmt.Printf("Remainder of %d %% %d: %d\n", intNum, anotherIntNum, remainder) // Output: Remainder of 10 % 4: 2

    // Demonstrating compound assignment operators
    var totalScore float64 = 0.0
    totalScore += sum // Adding the sum to totalScore
    fmt.Printf("Total Score after addition: %.2f\n", totalScore) // Output: Total Score after addition: 13.50

    totalScore *= 2   // Doubling the totalScore
    fmt.Printf("Total Score after multiplication: %.2f\n", totalScore) // Output: Total Score after multiplication: 27.00
}

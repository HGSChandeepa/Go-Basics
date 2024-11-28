package main

import "fmt"

func main() {
    // Declare and initialize variables
    a := 10
    b := 3

    // Perform arithmetic operations
    sum := a + b
    difference := a - b
    product := a * b
    quotient := a / b
    remainder := a % b

    // Print results
    fmt.Println("Sum:", sum)                // Output: Sum: 13
    fmt.Println("Difference:", difference)  // Output: Difference: 7
    fmt.Println("Product:", product)        // Output: Product: 30
    fmt.Println("Quotient:", quotient)      // Output: Quotient: 3
    fmt.Println("Remainder:", remainder)    // Output: Remainder: 1

    // Demonstrating compound assignment operators
    c := 5
    c += 2 // c = c + 2
    fmt.Println("After addition assignment, c:", c) // Output: c: 7

    c -= 1 // c = c - 1
    fmt.Println("After subtraction assignment, c:", c) // Output: c: 6

    c *= 3 // c = c * 3
    fmt.Println("After multiplication assignment, c:", c) // Output: c: 18

    c /= 2 // c = c / 2
    fmt.Println("After division assignment, c:", c) // Output: c: 9

    c %= 4 // c = c % 4
    fmt.Println("After modulo assignment, c:", c) // Output: c: 1



	// Increment and decrement operators --------------------------------------------

	d := 10
	d++ // Increment d by 1
	fmt.Println("After increment, d:", d) // Output: d: 11

	d-- // Decrement d by 1
	fmt.Println("After decrement, d:", d) // Output: d: 10


	// Bitwise operators -----------------------------------------------------------


	// Bitwise AND
	fmt.Println("Bitwise AND:", 5 & 3) // Output: 1

	// Bitwise OR
	fmt.Println("Bitwise OR:", 5 | 3) // Output: 7

	// Bitwise XOR
	fmt.Println("Bitwise XOR:", 5 ^ 3) // Output: 6

	// Bitwise AND NOT
	fmt.Println("Bitwise AND NOT:", 5 &^ 3) // Output: 4

	// Bitwise left shift
	fmt.Println("Bitwise left shift:", 5 << 1) // Output: 10

	// Bitwise right shift
	fmt.Println("Bitwise right shift:", 5 >> 1) // Output: 2

	// Bitwise NOT
	fmt.Println("Bitwise NOT:", ^5) // Output: -6
}

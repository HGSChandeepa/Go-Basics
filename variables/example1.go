package main

import "fmt"

const pi = 3.14 // Constant declaration

func main() {
    // Variable declarations
    var radius float64 = 5.0
    var area float64

    // Using pointers
    pointerToRadius := &radius // Pointer to radius
    *pointerToRadius = 10.0 // Change value via pointer

    // Calculate area using constant pi
    area = pi * (*pointerToRadius) * (*pointerToRadius)

    fmt.Printf("The area of the circle with radius %.2f is %.2f\n", *pointerToRadius, area) // %.2f formats the float to 2 decimal places

    // Short variable declaration for multiple variables
    length, width := 10, 5
    fmt.Printf("Length: %d, Width: %d\n", length, width)
}

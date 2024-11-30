package main

import "fmt"

// Function to increment a value using a pointer
func increment(value *int) {
    *value++
}

// Function to set a value to zero using a pointer
func setToZero(value *int) {
    *value = 0 
}

func main() {
    num := 10

    fmt.Println("Original value:", num)

    increment(&num)
    fmt.Println("After incrementing:", num)

    setToZero(&num) 
    fmt.Println("After setting to zero:", num)
}

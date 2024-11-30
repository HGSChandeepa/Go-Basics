package main

import "fmt"

func main() {
    // Declare an integer variable
    var num int = 42

    // Declare a pointer to an integer
    var ptr *int

    // Assign the address of num to ptr
    ptr = &num

    // Print the original value of num
    fmt.Println("Original value of num:", num) // Output: 42

    // Print the value of ptr (address of num)
    fmt.Println("Address of num:", ptr) // Output: Address of num

    // Dereference ptr to get the value it points to
    fmt.Println("Value pointed to by ptr:", *ptr) // Output: 42

    // Modify the value of num using the pointer
    *ptr = 100

    // Print the modified value of num
    fmt.Println("Modified value of num:", num) // Output: 100
}

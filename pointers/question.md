# Question: Dynamic Array Management with Pointers

## Problem Statement

You are tasked with creating a simple program that manages a dynamic array of integers using pointers. Your program should perform the following tasks:

1. Define a struct called `DynamicArray` that contains:
    - A slice of integers to hold the array elements.
    - An integer to track the current size of the array.
2. Implement the following methods for the `DynamicArray` struct:
    - `Add`: A method that takes an integer as an argument and adds it to the dynamic array. If the current size equals the length of the underlying slice, double the size of the slice before adding the new element.
    - `Print`: A method that prints all elements in the dynamic array.
    - `Remove`: A method that takes an index as an argument and removes the element at that index, shifting subsequent elements left.
3. In the `main` function:
    - Create an instance of `DynamicArray`.
    - Add several integers to the dynamic array.
    - Print the contents of the dynamic array.
    - Remove an element from a specified index and print the updated contents.

## Requirements

- Ensure that your program compiles and runs without errors.
- Handle edge cases (e.g., removing an element from an empty array or invalid index).
- Use appropriate variable names and comments to explain your code.

## Sample Output

```
Current Array: [1 2 3 4 5]
Removing element at index 2...
Updated Array: [1 2 4 5]
```

## Solution

```go
package main

import "fmt"

// Define the DynamicArray struct
type DynamicArray struct {
    elements []int // Slice to hold array elements
    size     int   // Current size of the array
}

// Method to add an element to the dynamic array
func (da *DynamicArray) Add(value int) {
    // Check if we need to resize
    if da.size == len(da.elements) {
        // Double the size of the slice
        newSlice := make([]int, da.size*2)
        copy(newSlice, da.elements) // Copy old elements to new slice
        da.elements = newSlice       // Update reference to new slice
    }
    // Add new value and increment size
    da.elements[da.size] = value
    da.size++
}

// Method to print current elements in the dynamic array
func (da *DynamicArray) Print() {
    fmt.Print("Current Array: ")
    for i := 0; i < da.size; i++ {
        fmt.Print(da.elements[i], " ")
    }
    fmt.Println()
}

// Method to remove an element at a specific index
func (da *DynamicArray) Remove(index int) {
    if index < 0 || index >= da.size {
        fmt.Println("Index out of bounds")
        return
    }
    
    // Shift elements left to fill gap
    for i := index; i < da.size-1; i++ {
        da.elements[i] = da.elements[i+1]
    }
    
    // Decrease size
    da.size--
}

func main() {
    // Create an instance of DynamicArray with initial capacity
    da := &DynamicArray{
        elements: make([]int, 0, 5), // Initial capacity of 5
        size:     0,
    }

    // Add elements to the dynamic array
    for i := 1; i <= 5; i++ {
        da.Add(i)
    }

    // Print current contents of the dynamic array
    da.Print()

    // Remove an element at index 2 (third element)
    fmt.Println("Removing element at index 2...")
    da.Remove(2)

    // Print updated contents of the dynamic array
    da.Print()
}
```

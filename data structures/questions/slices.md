## Problem Statement

You are tasked with writing a function to remove an element from a slice based on its index. Your program should perform the following tasks:

1. Create a slice of integers containing the following values: `[10, 20, 30, 40, 50, 60]`.
2. Write a function called `removeElement` that takes two parameters:
    - A slice of integers.
    - An index of the element to be removed.
3. The function should remove the element at the specified index and return the modified slice.
4. In the main function, call the `removeElement` function to remove an element from the slice and print the resulting slice.

### Requirements

- Ensure that your program handles cases where the index is out of bounds gracefully.
- Use appropriate variable names and comments to explain your code.

### Sample Output

```
Original Slice: [10 20 30 40 50 60]
Slice after removing element at index 3: [10 20 30 50 60]
```

### Solution

```go
package main

import "fmt"

// Function to remove an element from a slice by index
func removeElement(slice []int, index int) []int {
    // Check if the index is out of bounds
    if index < 0 || index >= len(slice) {
        fmt.Println("Index out of bounds")
        return slice // Return original slice if index is invalid
    }

    // Remove the element at the specified index
    return append(slice[:index], slice[index+1:]...) // Append elements before and after the index
}

func main() {
    // Create and initialize a slice of integers
    numbers := []int{10, 20, 30, 40, 50, 60}
    fmt.Println("Original Slice:", numbers)

    // Specify the index of the element to remove
    indexToRemove := 3

    // Call removeElement function
    modifiedSlice := removeElement(numbers, indexToRemove)

    // Print the modified slice
    fmt.Println("Slice after removing element at index", indexToRemove, ":", modifiedSlice)
}
```

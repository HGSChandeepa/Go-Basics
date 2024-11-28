/*
2. Slices
Definition: A slice is a more flexible and powerful data structure than an array. It can grow and shrink in size, and it does not require a predefined length.
*/

package main

import "fmt"

func main() {
    // Declare and initialize a slice
    slice := []int{1, 2, 3, 4, 5}
    
    // Append elements to the slice
    slice = append(slice, 6)
    
    fmt.Println("Slice after append:", slice) // Output: Slice after append: [1 2 3 4 5 6]

	// Length of a slice
	fmt.Println("Length of slice:", len(slice)) // Output: Length of slice: 6

	// Iterating over a slice
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Adding elements to a slice
	slice = append(slice, 7, 8, 9)
	fmt.Println("Slice after adding elements:", slice) // Output: Slice after adding elements: [1 2 3 4 5 6 7 8 9]

	// Removing elements from a slice
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("Slice after removing element at index 2:", slice) // Output: Slice after removing element at index 2: [1 2 4 5 6 7 8 9]

	// Updating elements in a slice
	slice[0] = 10
	fmt.Println("Modified first element:", slice[0]) // Output: Modified first element: 10

}

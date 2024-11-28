/*

1. Arrays
Definition: An array is a fixed-size collection of elements of the same type. Once defined, the size of an array cannot change.

*/

package main

import "fmt"

func main() {
    // Declare and initialize an array
    var numbers [5]int = [5]int{1, 2, 3, 4, 5}
    
    // Accessing array elements
    fmt.Println("First number:", numbers[0]) // Output: First number: 1
    
    // Modifying an element
    numbers[0] = 10
    fmt.Println("Modified first number:", numbers[0]) // Output: Modified first number: 10

	// Length of an array
	fmt.Println("Length of numbers array:", len(numbers)) // Output: Length of numbers array: 5

	// Iterating over an array
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Adding elements to an array : we can't add elements to an array once it's defined

	// Removing elements from an array : we can't remove elements from an array once it's defined

	// updating elements in an array
	numbers[0] = 10
	fmt.Println("Modified first number:", numbers[0]) // Output: Modified first number: 10

	// Declare and initialize an array without specifying the size
	var numbers2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Length of numbers2 array:", len(numbers2)) // Output: Length of numbers2 array: 5

	// Multidimensional arrays
	var matrix [2][2]int = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix) // Output: [[1 2] [3 4]]

	// Iterating over a multidimensional array
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}

	
}

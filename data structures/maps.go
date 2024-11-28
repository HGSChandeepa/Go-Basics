/*
3. Maps
Definition: A map is an unordered collection of key-value pairs. It allows you to associate values with unique keys for fast lookups.

*/

package main

import "fmt"

func main() {
    // Declare and initialize a map
    studentAges := map[string]int{
        "Alice": 20,
        "Bob":   22,
        "Charlie": 23,
    }
    
    // Accessing map values using keys
    fmt.Println("Alice's age:", studentAges["Alice"]) // Output: Alice's age: 20
    
    // Adding a new key-value pair
    studentAges["David"] = 21
    
    fmt.Println("Updated map:", studentAges) // Output: Updated map: map[Alice:20 Bob:22 Charlie:23 David:21]

	// Length of a map
	fmt.Println("Length of studentAges map:", len(studentAges)) // Output: Length of studentAges map: 4

	// Iterating over a map
	for key, value := range studentAges {
		fmt.Println(key, value)
	}

	// Removing a key-value pair
	delete(studentAges, "Bob")
	fmt.Println("Map after deleting Bob:", studentAges) // Output: Map after deleting Bob: map[Alice:20 Charlie:23 David:21]

	// Updating a value in a map
	studentAges["Alice"] = 21
	fmt.Println("Updated Alice's age:", studentAges["Alice"]) // Output: Updated Alice's age: 21

	// Checking if a key exists in a map
	age, ok := studentAges["Alice"] // ok is true if the key exists in the map and false otherwise
	if ok {
		fmt.Println("Alice's age is", age)
	} else {
		fmt.Println("Alice's age is not available")
	}
}

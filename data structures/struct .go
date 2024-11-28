package main

import "fmt"

// Define a struct type for Person
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    // Create an instance of Person
    person := Person{FirstName: "John", LastName: "Doe", Age: 30}
    
    // Accessing struct fields
    fmt.Println("Person:", person.FirstName, person.LastName, "is", person.Age, "years old.")

	// Updating struct fields
	person.Age = 31

	fmt.Println("Updated age:", person.Age) // Output: Updated age: 31

	// Declare and initialize a struct without specifying field names
	person2 := Person{"Alice", "Smith", 25}
	fmt.Println("Person2:", person2.FirstName, person2.LastName, "is", person2.Age, "years old.")

	
}

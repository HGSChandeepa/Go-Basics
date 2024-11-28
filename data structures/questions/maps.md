### Question: Managing a Contact List with Maps

#### Problem Statement
You are tasked with creating a simple contact management program using maps in Go. Your program should perform the following tasks:

1. Create a map called `contacts` where the key is a string representing the contact name and the value is a string representing the phone number.
2. Initialize the map with the following contacts:
    - "Alice": "123-456-7890"
    - "Bob": "987-654-3210"
    - "Charlie": "555-555-5555"
3. Implement the following functionalities:
    - Print all contacts in the format: `Name: Phone Number`
    - Add a new contact to the map (e.g., `"David": "111-222-3333"`).
    - Update an existing contact's phone number (e.g., change Bob's number to `"222-333-4444"`).
    - Remove a contact from the map (e.g., remove Charlie).
4. Print the updated list of contacts after each operation.

#### Requirements
- Ensure that your program handles cases where you try to update or remove contacts that do not exist.
- Use appropriate variable names and comments to explain your code.

#### Sample Output
```
Contacts:
Alice: 123-456-7890
Bob: 987-654-3210
Charlie: 555-555-5555

Adding new contact David: 111-222-3333
Updating Bob's phone number to 222-333-4444
Removing Charlie from contacts.

Updated Contacts:
Alice: 123-456-7890
Bob: 222-333-4444
David: 111-222-3333
```

#### Solution
```go
package main

import "fmt"

func main() {
    // Step 1: Create and initialize the contacts map
    contacts := map[string]string{
        "Alice": "123-456-7890",
        "Bob": "987-654-3210",
        "Charlie": "555-555-5555",
    }

    // Step 2: Print all contacts
    fmt.Println("Contacts:")
    printContacts(contacts)

    // Step 3: Add a new contact
    fmt.Println("\nAdding new contact David: 111-222-3333")
    contacts["David"] = "111-222-3333"

    // Update Bob's phone number
    fmt.Println("Updating Bob's phone number to 222-333-4444")
    contacts["Bob"] = "222-333-4444"

    // Remove Charlie from contacts
    fmt.Println("Removing Charlie from contacts.")
    delete(contacts, "Charlie")

    // Print updated contacts
    fmt.Println("\nUpdated Contacts:")
    printContacts(contacts)
}

// Function to print all contacts in the format Name: Phone Number
func printContacts(contacts map[string]string) {
    for name, phone := range contacts {
        fmt.Printf("%s: %s\n", name, phone)
    }
}
```

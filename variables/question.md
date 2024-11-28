# Exercise: Simple Shopping Cart

## Problem Statement

You are tasked with creating a simple shopping cart program that performs the following tasks:

- Declare and initialize variables to store the following information:
  - The name of the product (string)
  - The price of the product (float64)
  - The quantity of the product (int)
  - A boolean indicating if the product is in stock
- Calculate the total cost of the items in the cart (price \* quantity) and print it.
- If the product is not in stock, print a message indicating that the item is unavailable.
- Use appropriate variable names and comments to explain your code.

## Requirements

- Ensure that your program compiles and runs without errors.
- Use basic variable types (string, int, float64, bool) for this exercise.

## Sample Output

```
Product Name: Laptop
Price: 999.99
Quantity: 2
In Stock: true
Total Cost: 1999.98

Product Name: Headphones
Price: 199.99
Quantity: 0
In Stock: false
The item is unavailable.
```

## Solution

```go
package main

import "fmt"

func main() {
        // Declare and initialize variables for the first product
        var productName1 string = "Laptop"
        var productPrice1 float64 = 999.99
        var productQuantity1 int = 2
        var isInStock1 bool = true

        // Calculate total cost for the first product
        totalCost1 := productPrice1 * float64(productQuantity1)

        // Print details of the first product
        fmt.Printf("Product Name: %s\n", productName1)
        fmt.Printf("Price: %.2f\n", productPrice1)
        fmt.Printf("Quantity: %d\n", productQuantity1)
        fmt.Printf("In Stock: %t\n", isInStock1)

        if isInStock1 {
                fmt.Printf("Total Cost: %.2f\n\n", totalCost1)
        } else {
                fmt.Println("The item is unavailable.\n")
        }

        // Declare and initialize variables for the second product
        var productName2 string = "Headphones"
        var productPrice2 float64 = 199.99
        var productQuantity2 int = 0
        var isInStock2 bool = false

        // Calculate total cost for the second product
        totalCost2 := productPrice2 * float64(productQuantity2)

        // Print details of the second product
        fmt.Printf("Product Name: %s\n", productName2)
        fmt.Printf("Price: %.2f\n", productPrice2)
        fmt.Printf("Quantity: %d\n", productQuantity2)
        fmt.Printf("In Stock: %t\n", isInStock2)

        if isInStock2 {
                fmt.Printf("Total Cost: %.2f\n", totalCost2)
        } else {
                fmt.Println("The item is unavailable.")
        }
}
```

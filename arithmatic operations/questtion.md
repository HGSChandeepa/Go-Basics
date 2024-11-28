### Question: Mixed Type Arithmetic Operations

#### Problem Statement
You are tasked with creating a simple program that calculates the total cost of items in a shopping cart. The program should perform the following tasks:

1. Declare and initialize the following variables:
    - `itemPrice` (float64) representing the price of a single item.
    - `itemQuantity` (int) representing the number of items purchased.
    - `discountPercentage` (float64) representing the discount on the total cost as a percentage.
2. Calculate the total cost before discount using the formula:
    - `Total Cost = itemPrice × itemQuantity`
3. Calculate the total discount using the formula:
    - `Total Discount = (Total Cost × discountPercentage / 100)`
4. Calculate the final cost after applying the discount using the formula:
    - `Final Cost = Total Cost − Total Discount`
5. Print out the following information:
    - Total Cost
    - Total Discount
    - Final Cost

#### Requirements
- Use appropriate variable names and comments to explain your code.
- Ensure that your program compiles and runs without errors.

#### Sample Output
```
Item Price: 49.99
Item Quantity: 3
Discount Percentage: 10.0

Total Cost: 149.97
Total Discount: 14.99
Final Cost: 134.98
```

#### Solution
```go
package main

import "fmt"

func main() {
    // Declare and initialize variables
    var itemPrice float64 = 49.99      // Price of a single item
    var itemQuantity int = 3            // Number of items purchased
    var discountPercentage float64 = 10.0 // Discount percentage

    // Calculate total cost before discount
    totalCost := itemPrice * float64(itemQuantity) // Convert itemQuantity to float64 for multiplication

    // Calculate total discount
    totalDiscount := totalCost * (discountPercentage / 100) // Calculate discount

    // Calculate final cost after applying the discount
    finalCost := totalCost - totalDiscount

    // Print out the results
    fmt.Printf("Item Price: %.2f\n", itemPrice)
    fmt.Printf("Item Quantity: %d\n", itemQuantity)
    fmt.Printf("Discount Percentage: %.1f\n\n", discountPercentage)

    fmt.Printf("Total Cost: %.2f\n", totalCost)
    fmt.Printf("Total Discount: %.2f\n", totalDiscount)
    fmt.Printf("Final Cost: %.2f\n", finalCost)
}
```

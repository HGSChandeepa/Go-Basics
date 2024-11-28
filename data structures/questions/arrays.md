## Question: Array Manipulation

### Problem Statement

You are tasked with creating a program that manages a list of temperatures recorded over a week. Your program should perform the following tasks:

1. Declare and initialize an array called `temperatures` to hold the daily temperatures for 7 days (use `float64` for temperature values).
2. Fill the array with the following temperatures (in degrees Celsius):
    - Day 1: 22.5
    - Day 2: 24.0
    - Day 3: 19.5
    - Day 4: 21.0
    - Day 5: 23.5
    - Day 6: 20.0
    - Day 7: 18.5
3. Calculate and print:
    - The average temperature for the week.
    - The highest temperature recorded.
    - The lowest temperature recorded.
4. Print the temperatures in a formatted manner, showing each day's temperature.

### Requirements

- Use appropriate variable names and comments to explain your code.
- Ensure that your program compiles and runs without errors.

### Sample Output

```
Temperatures for the week:
Day 1: 22.5°C
Day 2: 24.0°C
Day 3: 19.5°C
Day 4: 21.0°C
Day 5: 23.5°C
Day 6: 20.0°C
Day 7: 18.5°C

Average Temperature: 21.14°C
Highest Temperature: 24.0°C
Lowest Temperature: 18.5°C
```

### Solution

```go
package main

import "fmt"

func main() {
    // Declare and initialize the temperatures array
    temperatures := [7]float64{
        22.5, // Day 1
        24.0, // Day 2
        19.5, // Day 3
        21.0, // Day 4
        23.5, // Day 5
        20.0, // Day 6
        18.5, // Day 7
    }

    // Variables to calculate average, highest, and lowest temperatures
    var total float64 = 0
    highest := temperatures[0]
    lowest := temperatures[0]

    // Print the temperatures for each day and calculate total, highest, and lowest
    fmt.Println("Temperatures for the week:")
    for i := 0; i < len(temperatures); i++ {
        fmt.Printf("Day %d: %.1f°C\n", i+1, temperatures[i])
        total += temperatures[i] // Accumulate total temperature

        // Check for highest temperature
        if temperatures[i] > highest {
            highest = temperatures[i]
        }

        // Check for lowest temperature
        if temperatures[i] < lowest {
            lowest = temperatures[i]
        }
    }

    // Calculate average temperature
    average := total / float64(len(temperatures))

    // Print the results
    fmt.Printf("\nAverage Temperature: %.2f°C\n", average)
    fmt.Printf("Highest Temperature: %.1f°C\n", highest)
    fmt.Printf("Lowest Temperature: %.1f°C\n", lowest)
}
```

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Create an empty integer slice of size 3
	sliceVar := make([]int, 0, 3)

	for {
		var input string
		fmt.Print("Enter an integer (X to quit): ")
		fmt.Scanln(&input)

		if input == "X" || input == "x" {
			break
		}

		// Convert string to integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		// Append the integer to the slice
		sliceVar = append(sliceVar, num)

		// Sort the slice
		sort.Ints(sliceVar)

		// Print the sorted slice
		fmt.Println("Sorted slice:", sliceVar)
	}
}

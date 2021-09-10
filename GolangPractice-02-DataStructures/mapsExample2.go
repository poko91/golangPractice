package main

import "fmt"

func main() {

	// Create a map of person name and age

	m := map[string]int{
		"Ted":        14,
		"Maria":      30,
		"Roberto":    57,
		"Alessandro": 46,
		"Agnes":      78,
		"Richard":    12,
		"Alison":     24,
		"Dennis":     18,
	}

	// range on map iterates over key/value pairs.
	fmt.Println("Even or Odd numbers:")
	for i, v := range m {
		// using if statement for checking the condition
		if v%2 == 0 {
			fmt.Println(i, "is", v, "years old -- even number.")
			// code executes if condition is true
		} else {
			fmt.Println(i, "is", v, "years old -- odd number.")
			// code executes if condition is false
		}

	}

	fmt.Println("\nCheck whether person is eligible to drive or not:")

	for i, v := range m {
		// switch on values if condition is true
		switch {
		case v >= 18:
			//if value satisfies the condition
			//print the following
			fmt.Println(i, "-- eligible")
		case v <= 18:
			//if value satisfies the condition
			//print the following
			fmt.Println(i, "-- not eligible")
		}
	}
}

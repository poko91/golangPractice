package main

import "fmt"

func main() {
	// Declare and initialise a map of student names and grades recieved
	grades := map[string]string{
		"Lucia":  "A",
		"Paul":   "D",
		"Merry":  "B",
		"Jane":   "C",
		"George": "B",
		"Jack":   "C",
		"Nelson": "A",
		"Grace":  "C",
		"Nadia":  "B",
	}

	// range over the map to check if the grades are good or bad using if condition
	for name, grade := range grades {
		if grade < "C" {
			fmt.Println(name, "--", grade, "- Good Grade")
		} else {
			fmt.Println(name, "--", grade, "- Bad Grade. Need to take a test again")
		}
	}

	// check if a key exists in a map
	if _, ok := grades["Paul"]; ok {
		fmt.Println("\nStudent grade available")
	}

	// if key does not exist the default value is zero
	if _, ok := grades["Ashley"]; ok {
		fmt.Print("\nStudent grade not available")
	}
}

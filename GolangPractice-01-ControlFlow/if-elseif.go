package main

import "fmt"

func main() {
	m := map[string]int{
		"Ted":        14,
		"Maria":      30,
		"Roberto":    57,
		"Alessandro": 46,
	}

	m["Agnes"] = 78
	m["Richard"] = 12
	m["Alison"] = 24
	m["Dennis"] = 18
	//add more names

	// range on map iterates over key/value pairs.
	for i, v := range m {
		// using if statement for checking the condition
		if v%2 == 0 {
			fmt.Println(i, "is", v, "years old which is a even number.")
			// code executes if condition is true
		} else {
			fmt.Println(i, "is", v, "years old which is a odd number.")
			// code executes if condition is false
		}
	}
}

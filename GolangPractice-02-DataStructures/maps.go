package main

import "fmt"

func main() {

	// Declare an empty map
	var person = map[string]int{}
	fmt.Println("1. Empty map created having string as key-type and int as value-type:")
	fmt.Println(person)

	// Map declaration using make function
	m := make(map[string]int)

	// Adding key-value pairs to the map
	m["Agnes"] = 78
	m["Richard"] = 12
	m["Alison"] = 24
	m["Dennis"] = 18
	fmt.Println("\n2. Adding items (key-value pairs) to the map:")
	fmt.Println(m)

	// You can also declare and initialize a new map in the same line with this syntax: (map[key-type]val-type){}
	nm := map[string]int{
		"Ted":        14,
		"Maria":      30,
		"Roberto":    57,
		"Alessandro": 46,
	}
	fmt.Println("\n3. Declare and initialize a new map in the same line:")
	fmt.Println(nm)

	// The builtin delete removes key/value pairs from a map.
	fmt.Println("\n4. Deleting items from a map:")
	fmt.Println("Before deleting a value:", m)

	delete(m, "Agnes")
	fmt.Println("After deleting a value:", m)

	// Accessing Items of a map
	fmt.Println("\n5. Accessing specific items from a map:")
	fmt.Println(nm["Maria"])

	//Adding elements to an existing map
	fmt.Println("\n6. Adding items to a map:")

	m["Rocky"] = 30 // Add element
	m["Joseph"] = 40
	fmt.Println(m)

	// Changing values of existing map items
	fmt.Println("\n7. Changing values in a map:")
	m["Richard"] = 84
	m["Dennis"] = 12
	fmt.Println(m) // Edit item

	// Iterate over a Map
	fmt.Println("\n8. Iterating over a map:")
	for key, value := range m {
		fmt.Println(key, "--", value)
	}

}

package main

import "fmt"

func main() {
	// To create an empty map use syntax - make(map[key-type]val-type)
	m := make(map[string]int)

	// To add more key-value pairs to the map use syntax: m[key] = value
	m["Agnes"] = 78
	m["Richard"] = 12
	m["Alison"] = 24
	m["Dennis"] = 18
	fmt.Println(m)

	// You can also declare and initialize a new map in the same line with this syntax: (map[key-type]val-type){}
	nm := map[string]int{
		"Ted":        14,
		"Maria":      30,
		"Roberto":    57,
		"Alessandro": 46,
	}
	fmt.Println(nm)

	// The builtin delete removes key/value pairs from a map.
	delete(m, "Agnes")
	delete(nm, "Ted")

	fmt.Println("After deleting a value:", m)
	fmt.Println("After deleting a value:", nm)

}

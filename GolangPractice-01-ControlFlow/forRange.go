package main

import "fmt"

func main() {

	m := map[string]int{
		"Dishwasher":      28000,
		"Washing Machine": 22690,
		"Air Conditioner": 34000,
		"Water Purifier":  8000,
		"Microwave Oven":  11399,
	}

	var total int //declare variable to be used

	for k, v := range m {
		total += v // sum = sum + v
		fmt.Println("Cost of", k, "is: Rs.", v)
	}
	fmt.Println("The total cost of all items is: Rs.", total)

}

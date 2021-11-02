package main

import "fmt"

func main() {
	fmt.Println("26 + 19 = ", mySum(26, 19))
	fmt.Println("81 + 52 = ", mySum(81, 52))
	fmt.Println("78 + 13 = ", mySum(78, 13))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

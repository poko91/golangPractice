package main

import "fmt"

func condition(n int) bool {
	return n > 3 && n < 9
}

func filter(numbers []int) []int {
	xs := []int{}
	for _, n := range numbers {
		if condition(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(xs)
}

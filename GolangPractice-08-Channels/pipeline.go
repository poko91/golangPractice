package main

import "fmt"

func main() {
	// c := gen(6, 8, 9, 3, 10, 5, 78, 32, 17)
	// out := sq(c)
	// for n := range out {
	// 	fmt.Println(n)
	// }

	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

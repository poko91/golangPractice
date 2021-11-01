package main

import "fmt"

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	total := 1
	go func() {
		for i := n; i > 0; i-- {
			total *= n
		}
		out <- total
		close(out)
	}()
	return out
}

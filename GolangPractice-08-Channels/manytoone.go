package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Wait()
		close(c)
	}()

	go func() {
		wg.Add(1)
		for i := 10; i < 20; i++ {
			c <- i
		}
		wg.Wait()
		close(c)
	}()

	go func() {
		wg.Done()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

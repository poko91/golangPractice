package main

import (
	"fmt"
	"sync"
)

// many function writing to the same channel
func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	//func 01
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()
	
	//func 02
	go func() {
		for i := 10; i < 20; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

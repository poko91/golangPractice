package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	d := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				c <- i
			}
			if i%2 != 0 {
				d <- i
			}
		}
	}()
	go func() {
		for {
			fmt.Println("From channel c:", <-c)
		}
	}()

	go func() {
		for {
			fmt.Println("From channel d:", <-d)
		}
	}()

	time.Sleep(time.Second)
}

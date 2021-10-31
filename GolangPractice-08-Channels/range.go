package main

<<<<<<< HEAD
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

	go func(){
		wg.Done()
		close(c)
	}

=======
import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

>>>>>>> e00d3f0b091a0350fbff80433171d9ecfc077c8f
	for n := range c {
		fmt.Println(n)
	}
}

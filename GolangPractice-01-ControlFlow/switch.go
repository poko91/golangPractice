package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		switch {
		case (i%2 == 0):
			fmt.Println(i, "is a even number")
		case (i%2 != 0):
			fmt.Println(i, "is a odd number")
		default:
			fmt.Println("This is not a number")
		}
	}
}

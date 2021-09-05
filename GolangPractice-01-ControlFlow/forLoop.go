package main

import "fmt"

func main() {

	i := 0
	for {
		i++
		if 1%2 == 0 {
			continue
		}

		fmt.Println(i)
		if i >= 50 {
			break
		}
		i++
	}
}

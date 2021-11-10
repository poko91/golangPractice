package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Error happened", err)
	}
}

// Output :
// Error happened open log.txt: no such file or directory

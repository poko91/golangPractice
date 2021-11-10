package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("log.txt")
	if err != nil {
		log.Println("Error happened", err)
	}
}

// Output :
// 2021/11/10 14:06:36 Error happened open log.txt: no such file or directory
// the error is logged in on stdout with date and timestamp

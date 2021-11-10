package main

import (
	"os"
)

func main() {
	_, err := os.Open("log.txt")
	if err != nil {
		panic(err)
	}
}

// Output :
// 2021/11/10 14:26:26 Error happened open log.txt: no such file or directory
// exit status 1

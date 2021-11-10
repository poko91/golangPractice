package main

import (
	"fmt"
	"log"
	"os"
)

//use func init to initialise - create new file
//where the error log message can be stored.
func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("Error happened", err)
	}
}

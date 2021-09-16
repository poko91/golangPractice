package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type your name: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Name: %s", input)

	scanner2 := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nType the year you were born: ")
	scanner2.Scan()
	input2, _ := strconv.ParseInt(scanner2.Text(), 10, 64)
	fmt.Printf("Age: %d years old", 2021-input2)
	// fmt.Printf("\nYou will be %d years old at the end of 2021", 2021-input2)
}

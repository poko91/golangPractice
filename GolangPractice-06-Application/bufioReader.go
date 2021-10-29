package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// readLineFromStandardInput()
	// readKeyPress()
	readTextUntilDelimiter()
}

// ReadLine() from input
// func readLineFromStandardInput() {
// 	reader := bufio.NewReader(os.Stdin)
// 	text, _, _ := reader.ReadLine()
// 	fmt.Println(text)
// }

//
// func readKeyPress() {
// 	reader := bufio.NewReader(os.Stdin)
// 	char, _, _ := reader.ReadRune()

// 	switch char {
// 	case '\n':
// 		fmt.Println("You've pressed 'Enter'")
// 		break

// 	default:
// 		fmt.Printf("You've entered '%s'", string(char))
// 		break
// 	}
// }

func readTextUntilDelimiter() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('/')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	fmt.Println(text)
}

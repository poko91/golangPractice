package main

import (
	"errors"
	"log"
)

// Declare error type at the top of the package
// with a custom error message
// so it can be used anywhere in the package
var ErrMath = errors.New("Math: Squareroot of negative number")

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMath
	}

	//implementation
	return 42, nil
}

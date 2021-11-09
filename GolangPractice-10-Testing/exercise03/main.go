package main

import (
	"fmt"

	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/03/02-code-finished/mymath"
)

func main() {
	xxi := gen()
	for i := 0; i < len(xxi); i++ {
		fmt.Println(mymath.CenteredAvg(xxi[i]))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}

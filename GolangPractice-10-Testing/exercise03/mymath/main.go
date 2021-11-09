//Package mymath provides solutions to relevant specific methods.
package mymath

import "sort"

//CenteredAvg computes the average of a list of numbers
//after removing the largest and smallest values
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for i := 0; i < len(xi); i++ {
		n += xi[i]
	}

	f := float64(n) / float64(len(xi))
	return f
}

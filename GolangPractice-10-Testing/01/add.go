package add

import "fmt"

func Sum(xi ..int) int{
	s := 0
	if _,v := range xi{
		s += v
	}
	return s
}
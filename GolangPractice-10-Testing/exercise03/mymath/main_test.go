package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
		test{[]int{2, 4, 6, 8, 10, 12}, 7},
		test{[]int{10, 20, 40, 60, 80}, 40},
	}

	for i := 0; i < len(tests); i++ {
		f := CenteredAvg(tests[i].data)
		if f != tests[i].answer {
			t.Error("Want", tests[i].answer, "Got", f)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 6, 9}))
	// Output:
	// 4
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6})
	}
}

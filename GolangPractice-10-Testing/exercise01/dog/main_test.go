package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(10)
	if y != 70 {
		t.Error("wrong result, wanted 70")
	}
}

func TestYearsTwo(t *testing.T) {
	y2 := YearsTwo(20)
	if y2 != 140 {
		t.Error("wrong result, wanted 140")
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(Years(20))
	// Output:
	// 140
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYeaTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(20)
	}
}

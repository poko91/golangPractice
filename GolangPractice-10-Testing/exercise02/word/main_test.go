package word

import (
	"fmt"
	"testing"

	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/02-code-finished/quote"
)

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("Got", n, "Want", 3)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// Output:
	// 3
}
func ExampleUseCount() {
	fmt.Println(UseCount("one two three three three"))
	// Output:
	// map[one:1 three:3 two:1]
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

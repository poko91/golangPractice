package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(26, 19)
	if x != 45 {
		t.Error("Expected", 45, "Got", x)
	}
}

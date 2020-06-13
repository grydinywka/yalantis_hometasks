package main

import (
	"testing"

	HW "./homework3"
)

func TestAbs(t *testing.T) {
	square := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	exp := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	got := HW.SliceSnail(square)
	if len(got) != len(exp) {
		t.Errorf("Expected %d, received %d", exp, got)
	} else {
		for i := 0; i < len(got); i++ {
			if got[i] != exp[i] {
				t.Errorf("Expected %d, received %d", exp[i], got[i])
			}
		}
	}
}

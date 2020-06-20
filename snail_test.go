package main

import (
	"testing"

	HW "./homework3"
)

func checks(t *testing.T, got []int, exp []int, err error) {
	if err != nil {
		t.Logf(err.Error())
	}
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

func TestSnailSize6(t *testing.T) {
	square := [][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24}, {25, 26, 27, 28, 29, 30}, {31, 32, 33, 34, 35, 36}}
	exp := []int{1, 2, 3, 4, 5, 6, 12, 18, 24, 30, 36, 35, 34, 33, 32, 31, 25, 19, 13, 7, 8, 9, 10, 11, 17, 23, 29, 28, 27, 26, 20, 14, 15, 16, 22, 21}
	got, err := HW.SliceSnail(square)
	checks(t, got, exp, err)
}

func TestSnailSize3(t *testing.T) {
	square := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	exp := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	got, err := HW.SliceSnail(square)
	checks(t, got, exp, err)
}

func TestSnailPanicRowsNotEqualLen(t *testing.T) {
	// defer func() { recover() }()
	square := [][]int{{1, 2, 3}, {4, 5, 6, 100500}, {7, 8, 9}}
	got, err := HW.SliceSnail(square)
	checks(t, got, []int{}, err)
	// Never reaches here if `OtherFunctionThatPanics` panics.
	// t.Errorf("did not panic")
}

func TestSnailPanicWidthHeightNoEqual(t *testing.T) {
	// defer func() { recover() }()
	square := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	got, err := HW.SliceSnail(square)
	checks(t, got, []int{}, err)
	// Never reaches here if `OtherFunctionThatPanics` panics.
	// t.Errorf("did not panic")
}

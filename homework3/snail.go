package main

import (
	"fmt"
)

func SliceSnail(square [][]int) []int {
	line := []int{}

	// we get size of square
	width := len(square)
	height := 0
	fmt.Println(width, "width")
	for row, val := range square {
		fmt.Println(row, val)
		currentLen := len(val)
		if row != 0 && height != currentLen {
			panic("len of every of rows must be the same")
		}
		height = currentLen
	}
	if width != height {
		panic("Width and height are diff")
	}

	min, max := 0, width-1
	for ; min < max; min, max = min+1, max-1 {
		row := min
		for col := min; col < max; col++ {
			line = append(line, square[row][col])
		}
		col := max
		for row := min; row < max; row++ {
			line = append(line, square[row][col])
		}
		row = max
		for col := max; col > min; col-- {
			line = append(line, square[row][col])
		}
		col = min
		for row := max; row > min; row-- {
			line = append(line, square[row][col])
		}

	}
	if min == max {
		line = append(line, square[min][max])
	}

	return line
}

func main() {
	// make square
	size := 4
	square := make([][]int, size)

	for row, item := range square {
		for col := 0; col < size; col++ {
			item = append(item, size*(row+1)-size+col+1)
		}
		square[row] = item
	}
	// square := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(SliceSnail(square))
}

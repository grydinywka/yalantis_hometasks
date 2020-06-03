package main

import (
	"fmt"
)

func SliceSnail(size int) []int {
	// make square
	square := make([][]int, size)
	line := []int{}

	for row, item := range square {
		for col := 0; col < size; col++ {
			item = append(item, size*(row+1)-size+col+1)
		}
		square[row] = item
	}
	min, max := 0, size-1
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
	fmt.Println(SliceSnail(1))
    fmt.Println(SliceSnail(2))
    fmt.Println(SliceSnail(3))
    fmt.Println(SliceSnail(4))
    fmt.Println(SliceSnail(5))
}

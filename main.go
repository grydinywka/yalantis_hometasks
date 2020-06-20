package main

import (
	"fmt"
	"log"
	HW "./homework3"
	L "./lib"
)

func main() {
	var phrase string = "eat0 ! ((())) is e, G`oo0d dog!) `ert `rtjo 55 44 33 22 11"

	pLatin := L.ToPigLatin(&phrase)
	fmt.Println("pLatin = " + *pLatin)

	encoded := L.Encode(&phrase)
	fmt.Printf("Encode = %s\n", *encoded)

	decoded := L.Decode(&phrase)
	fmt.Printf("Encode = %s\n", *decoded)

	// make square
	size := 6
	square := make([][]int, size)

	for row, item := range square {
		for col := 0; col < size; col++ {
			item = append(item, size*(row+1)-size+col+1)
		}
		square[row] = item
	}
	// square := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	snail, err := HW.SliceSnail(square)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(snail)
}

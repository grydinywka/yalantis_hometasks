package main

import (
	"fmt"

	LT "./like_trello"
)

func main() {
	// var phrase string = "eat0 ! ((())) is e, G`oo0d dog!) `ert `rtjo 55 44 33 22 11"

	// pLatin := L.ToPigLatin(&phrase)
	// fmt.Println("pLatin = " + *pLatin)

	// encoded := L.Encode(&phrase)
	// fmt.Printf("Encode = %s\n", *encoded)

	// decoded := L.Decode(&phrase)
	// fmt.Printf("Encode = %s\n", *decoded)

	// // make square
	// size := 6
	// square := make([][]int, size)

	// for row, item := range square {
	// 	for col := 0; col < size; col++ {
	// 		item = append(item, size*(row+1)-size+col+1)
	// 	}
	// 	square[row] = item
	// }
	// // square := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// fmt.Println(HW.SliceSnail(square))

	fmt.Println("Trello")
	c := LT.Column{Name: "To do"}
	c2 := LT.Column{Name: "Doing"}
	c3 := LT.Column{Name: "To be Reviewed"}
	c4 := LT.Column{Name: "To be Tested"}
	b := LT.Board{Name: "The best board"}
	columns := []*LT.Column{&c, &c2, &c3, &c4}
	bm := LT.NewBoardManager(&b, columns...)
	fmt.Println(b)
	fmt.Println(bm, "bm")

	bm.CreateTask(&c, "Init task$$$$$", &LT.First)
	bm.CreateTask(&c, "Second task$$$$$", &LT.First)
	bm.CreateTask(&c, "Third task$$$$$", &LT.First)
	bm.CreateTask(&c, "Fourth task$$$$$", &LT.First)
	bm.CreateTask(&c, "Fifth task$$$$$", &LT.First)

	for _, t := range c.Tasks {
		fmt.Println(*t, "tasks - columns")
	}
	fmt.Println("--------------------")
	bm.MoveTop(&c, c.Tasks[4])
	for _, t := range c.Tasks {
		fmt.Println(*t, "tasks - columns")
	}
	bm.MoveToColumn(c.Tasks[4], &c2)
	// bm.MoveToColumn(c.Tasks[2], &c2)
	// fmt.Println(*c.Tasks[1], "tasks - columns")
	fmt.Println(len(c2.Tasks), c2.Tasks[0].Priority, "c2")
	// fmt.Println(len(c2.Tasks), c2.Tasks[1].Priority, "c2")

	fmt.Println(len(c.Tasks), "c")

	bm.AddComment(c.Tasks[0], "What is up?")
	fmt.Println(c.Tasks[0].Comments[0], "c")

	// bm2 := LT.NewBoardManager(&b)
	// fmt.Println(bm2.Columns[0], "bm2")
}

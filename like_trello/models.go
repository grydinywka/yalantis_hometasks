package like_trello

import "fmt"

// "fmt"

type Board struct {
	Name string `json:"name"`
	// Columns []Column `json:"columns"`
}

type Column struct {
	Name  string
	Board Board
	Tasks []Task
}

type Task struct {
	Name     string
	Status   int
	Priority int
	Comments []Comment
	Columns  []Column
}

type Comment struct {
	Question string
	Task     Task
}

type BoardManager struct {
	Board   Board
	Columns []*Column
}

func NewBoardManager(board *Board, columns ...*Column) *BoardManager {
	// bm := BoardManager{Board: board}
	for _, column := range columns {
		column.Board = *board
		fmt.Println("col", column)
	}
	return &BoardManager{Board: *board, Columns: columns}
}

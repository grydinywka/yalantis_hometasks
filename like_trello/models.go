package like_trello

import "fmt"

type Status int

var (
	First  Status = 1
	Second Status = 2
	Third  Status = 3
)

type Board struct {
	Name string `json:"name"`
	// Columns []Column `json:"columns"`
}

type Column struct {
	Name         string
	Board        *Board
	Tasks        []*Task
	PriorityLast int
}

type Task struct {
	Name     string
	Status   *Status
	Priority int
	Comments []*Comment
	Column   *Column
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
	for _, column := range columns {
		column.Board = board
		fmt.Println("col", column)
	}
	if len(columns) == 0 {
		columns = append(columns, &Column{Name: "Default", Board: board})
	}
	return &BoardManager{Board: *board, Columns: columns}
}

// CreateTask method add a task inside a column
func (bm *BoardManager) CreateTask(column *Column, name string, status *Status, comments ...*Comment) {
	task := &Task{Name: name, Status: status, Priority: bm.GetPriorityLast(column), Comments: comments}
	column.Tasks = append(column.Tasks, task)
	task.Column = column
}

func (bm *BoardManager) GetPriorityLast(column *Column) int {
	column.PriorityLast += 1
	return column.PriorityLast
}

func (bm *BoardManager) MoveTop(column *Column, task *Task) {
	smallest_priority := column.PriorityLast
	var topTask *Task = nil
	for _, taskObj := range column.Tasks {
		if taskObj.Priority < smallest_priority {
			smallest_priority = taskObj.Priority
			topTask = taskObj
		}
	}
	if topTask != nil {
		topTask.Priority, task.Priority = task.Priority, topTask.Priority
	}
}

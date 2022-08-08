package models

type Todo struct {
	Id        int
	Task      string
	Assignee  string
	Deadline  string
	Completed int
}

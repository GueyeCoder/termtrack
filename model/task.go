package model

type Task struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Deadline string `json:"deadline"`
	Priority string `json:"priority"`
}

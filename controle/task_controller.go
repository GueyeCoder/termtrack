package controller

import (
	"errors"

	"github.com/GueyeCoder/termtrack/model"
	"github.com/GueyeCoder/termtrack/storage"
)

type TaskController struct {
	tasks    []model.Task
	filePath string
}

func NewTaskController(filePath string) (*TaskController, error) {
	tasks, err := storage.LoadTasks(filePath)
	if err != nil {
		return nil, err
	}
	return &TaskController{
		tasks:    tasks,
		filePath: filePath,
	}, nil
}

func (c *TaskController) GetAll() []model.Task {
	return c.tasks
}

func (c *TaskController) Add(task model.Task) error {
	maxID := 0
	for _, t := range c.tasks {
		if t.Id > maxID {
			maxID = t.Id
		}
	}
	task.Id = maxID + 1

	c.tasks = append(c.tasks, task)
	return storage.SaveTasks(c.filePath, c.tasks)
}

func (c *TaskController) MarkDone(id int) error {
	for i, t := range c.tasks {
		if t.Id == id {
			c.tasks[i].Done = true
			return storage.SaveTasks(c.filePath, c.tasks)
		}
	}
	return errors.New("tâche non trouvée")
}

func (c *TaskController) Update(updatedTask model.Task) {
	for i, t := range c.tasks {
		if t.Title == updatedTask.Title {
			(c.tasks)[i] = updatedTask
			return
		}
	}
}

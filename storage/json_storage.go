package storage

import (
	"encoding/json"
	"log"
	"os"

	"github.com/GueyeCoder/termtrack/model"
)

func LoadTasks(filepath string) ([]model.Task, error) {
	var tasks []model.Task
	if _, err := os.Stat(filepath); err == nil {
		data, err := os.ReadFile(filepath)
		if err != nil {
			return []model.Task{}, nil
		}
		if err := json.Unmarshal(data, &tasks); err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

func SaveTasks(filePath string, tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatal("Error saving task! -", err)
	}
	return os.WriteFile(filePath, data, 0644)
}

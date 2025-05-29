package controle

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	tasks    = []Task{}
	taskFile = "tasks.json"
)

func loadTask() {
	if _, err := os.Stat(taskFile); err == nil {
		data, err := os.ReadFile(taskFile)
		if err != nil {
			log.Fatal("Error loading inventory file! -", err)
		}
		json.Unmarshal(data, &tasks)
	}
}

func saveTask() {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatal("Error saving task! -", err)
	}
	os.WriteFile(taskFile, data, 0644)
}

func deleteItem(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid item index")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	saveTask()
}

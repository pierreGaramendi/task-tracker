package tasks

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func ReadTasksFromFile(filename string) (*TaskList, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var taskList TaskList
	if err := json.Unmarshal(bytes, &taskList); err != nil {
		return nil, fmt.Errorf("error deserializing JSON: %v", err)
	}

	return &taskList, nil
}

func (tl *TaskList) PrintTasks(status TaskStatus) {
	fmt.Printf("Task List:\n")
	fmt.Printf("\n")
	for _, task := range tl.Tasks {
		if task.Status == status || status == "" {
			fmt.Printf("ID: %d\nDescription: %s\nState: %s\nCreated at: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt.Format("02 Jan 2006"))
			fmt.Printf("\n")
		}
	}
}

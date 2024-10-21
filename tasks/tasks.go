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

func (tl *TaskList) WriteTasksToFile(filename string) error {
	bytes, err := json.MarshalIndent(tl, "", "")
	if err != nil {
		return fmt.Errorf("error serializing task to JSON: %v", err)
	}
	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

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

func (tl *TaskList) AddTask(description string, status TaskStatus) int {
	newId := len(tl.Tasks) + 1
	newTask := Task{
		ID:          len(tl.Tasks) + 1,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tl.Tasks = append(tl.Tasks, newTask)
	return newId
}

func (tl *TaskList) DeleteTask(id int) error {
	index := -1
	for i, task := range tl.Tasks {
		if task.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("task with Id %d not found", id)
	}
	tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
	return nil
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

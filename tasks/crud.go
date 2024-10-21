package tasks

import (
	"fmt"
	"time"
)

func (tl *TaskList) PrintTasks(status TaskStatus) {
	if len(tl.Tasks) == 0 {
		fmt.Printf("Task list is empty.\n")
		return
	}
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

func (tl *TaskList) UpdateTask(id int, newDescription string) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Description = newDescription
			tl.Tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (tl *TaskList) UpdateStatusTask(id int, newStatus TaskStatus) error {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Status = newStatus
			tl.Tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func IsTaskNameValid(taskName string) bool {
	return taskName != ""
}

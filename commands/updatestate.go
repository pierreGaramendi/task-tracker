package commands

import (
	"flag"
	"fmt"
	"os"
	"task-tracker/tasks"
)

func HandleUpdateStateCommand(args []string) {
	updateStateCmd := flag.NewFlagSet("update-state", flag.ExitOnError)
	taskId := updateStateCmd.Int("id", 0, "Task Id")
	updateStateCmd.Parse(args[1:])
	if *taskId <= 0 {
		fmt.Println("A task id is required")
		updateStateCmd.PrintDefaults()
		os.Exit(1)
	}
	taskList, err := tasks.ReadTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	newStatus := GetStatusType(args[0])
	errUpdating := taskList.UpdateStatusTask(*taskId, newStatus)
	if errUpdating != nil {
		fmt.Printf("task with Id %d not found.\n", *taskId)
		return
	}
	if err := taskList.WriteTasksToFile("tasks.json"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Print("Task Updated successfully.\n")
	}
}

func GetStatusType(commandUpdate string) tasks.TaskStatus {
	if commandUpdate == "mark-in-progress" {
		return tasks.InProgress
	} else if commandUpdate == "mark-done" {
		return tasks.Done
	} else {
		return tasks.Todo
	}
}

package commands

import (
	"flag"
	"fmt"
	"os"
	"task-tracker/tasks"
)

func HandleUpdateCommand(args []string) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	taskId := updateCmd.Int("id", -1, "Task Id")
	desc := updateCmd.String("desc", "", "Task Description")
	updateCmd.Parse(args)
	if !tasks.IsTaskNameValid(*desc) {
		fmt.Println("A task description is required")
		updateCmd.PrintDefaults()
		os.Exit(1)
	}
	taskList, err := tasks.ReadTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	errUpdating := taskList.UpdateTask(*taskId, *desc)
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

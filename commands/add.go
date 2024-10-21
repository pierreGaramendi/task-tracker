package commands

import (
	"flag"
	"fmt"
	"os"
	"task-tracker/tasks"
)

func HandleAddCommand(args []string) {

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	taskName := addCmd.String("task", "", "Name of the task to add")
	addCmd.Parse(args)

	if !tasks.IsTaskNameValid(*taskName) {
		fmt.Println("A task name is required")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

	taskList, err := tasks.ReadTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newId := taskList.AddTask(*taskName, tasks.Todo)
	if err := taskList.WriteTasksToFile("tasks.json"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Task added successfully. (ID: %d)\n", newId)
	}
}

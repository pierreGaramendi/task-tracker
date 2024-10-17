package commands

import (
	"flag"
	"fmt"
	"log"
	"task-tracker/tasks"
)

func HandleListCommand(args []string) {
	addCmd := flag.NewFlagSet("list", flag.ExitOnError)
	status := addCmd.String("status", "", "Status of the task")
	addCmd.Parse(args)

	taskList, err := tasks.ReadTasksFromFile("tasks.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	if *status == "" {
		taskList.PrintTasks("")
	} else {
		taskStatus := tasks.TaskStatus(*status)
		if !taskStatus.IsValid() {
			fmt.Printf("Invalid status: %s. Valid statuses are: todo, in-progress, done.\n", *status)
			return
		}
		taskList.PrintTasks(taskStatus)
	}
}

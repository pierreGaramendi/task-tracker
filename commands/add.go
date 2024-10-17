package commands

import (
	"flag"
	"fmt"
	"os"
)

func HandleAddCommand(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	taskName := addCmd.String("task", "", "Name of the task to add")

	addCmd.Parse(args)

	if !IsTaskNameValid(*taskName) {
		fmt.Println("A task name is required")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Task Added: %s\n", *taskName)
}

func IsTaskNameValid(taskName string) bool {
	return taskName != ""
}

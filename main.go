package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("One subcommand is required: add")
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	taskName := addCmd.String("task-name", "", "Task name")

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *taskName == "" {
			fmt.Println("A task name is required")
			addCmd.PrintDefaults()
			os.Exit(1)
		}
		// TODO logic to add task
		fmt.Printf("Task added Id: %s\n", *taskName)
	default:
		fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
		fmt.Println("Available subcommands: list, add, update, delete, mark-in-progress, mark-done")
		os.Exit(1)
	}
}

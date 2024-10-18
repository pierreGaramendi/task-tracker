package main

import (
	"fmt"
	"os"
	"task-tracker/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("One subcommand is required")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		commands.HandleAddCommand(os.Args[2:])
	case "list":
		commands.HandleListCommand(os.Args[2:])
	case "delete":
		commands.HandleDeleteCommand(os.Args[2:])
	default:
		fmt.Printf("Unknown subcommand: %s\n", os.Args[1])
		fmt.Println("Available subcommands: list, add, update, delete, mark-in-progress, mark-done")
		os.Exit(1)
	}
}

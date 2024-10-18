package commands

import (
	"flag"
	"fmt"
	"task-tracker/tasks"
)

func HandleDeleteCommand(args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	taskId := deleteCmd.Int("id", -1, "Task id to delete")
	deleteCmd.Parse(args)

	taskList, err := tasks.ReadTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	errDelete := taskList.DeleteTask(*taskId)

	if errDelete != nil {
		fmt.Printf("task with Id %d not found.\n", *taskId)
		return
	}
	if err := taskList.WriteTasksToFile("tasks.json"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Print("Task deleted successfully.\n")
	}

}

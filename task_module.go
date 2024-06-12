package main

import (
	"fmt"
	"strings"

	"github.com/go-humble/locstor"
)

func HelpMessage() string {
	var message string
	message = "\nType 'clear' to clear the screen.\n"
	message += "Type 'add' to add a new task.'\n"
	message += "Type 'list' to list all tasks.\n"
	message += "Type 'done' to mark a task as done.\n"
	message = strings.ReplaceAll(message, "\n", "<br>")

	return message
}

func AddTask(task string) {
	count, err := locstor.Length()
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(count)
}

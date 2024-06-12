package main

func HelpMessage() string {
	var message string
	message = "Type 'clear' to clear the screen.\n"
	message += "Type 'add to add a new task.'\n"
	message += "Type 'list' to list all tasks.\n"
	message += "Type 'done' to mark a task as done.\n"

	return message
}

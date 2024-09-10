package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"task-tracker/app"
)

var (
	arg = os.Args[1]
)

func getId() int {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("error parsing id %v", err)
	}
	return id
}

func NewCmd(App *app.App) *Command {
	return &Command{App: *App}
}
func (c Command) Cmd() {
	switch arg {
	case "add":
		id := c.App.AddTask(os.Args[2])
		fmt.Printf("Task added successfully (ID: %d)", id)
	case "update":
		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("error parsing id %v", err)
		}
		c.App.UpdateTask(i, os.Args[3])
	case "delete":
		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("error parsing id %v", err)
		}
		c.App.DeleteTask(i)

	case "list":
		length := len(os.Args)
		if length == 3 {
			switch os.Args[2] {
			case "done":
				c.Table(c.App.ListDoneTasks)
			case "in-progress":
				c.Table(c.App.ListProgressTasks)
			case "todo":
				c.Table(c.App.ListToDoTasks)
			default:
				fmt.Println("Invalid command")
				fmt.Printf("Try:\n-'todo';\n-'in-progress;\n-'done';\n")
			}
		} else {
			c.Table(c.App.ListAllTasks)
		}
	case "mark-in-progress":
		id := getId()
		c.App.MarkInProgress(id)
	case "mark-done":
		id := getId()
		c.App.MarkDone(id)
	case "mark-todo":
		id := getId()
		c.App.MarkToDo(id)

	default:
		fmt.Println("\033[31m" + "Invalid command" + "\033[0m")
		fmt.Printf("List of commands:\n\033[34madd <«your_task»> \nupdate <task_id> <«your_task»>\ndelete <task_id>\nmark-in-progress <task_id>\nmark-done <task-id>\nmark-todo <task-id>\nlist #will list all tasks\nlist done \nlist todo\nlist in-progress\u001B[0m")
	}

}

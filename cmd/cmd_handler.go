package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-tracker/app"
	"task-tracker/lib/e"
	"task-tracker/models"
)

const (
	ErrArgumentCount = "argument count error"
	InvalidCmd       = "invalid command"
)

var (
	arg = os.Args[1]
)

func getId(arg string) (int, error) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewCmd(App *app.App) *Command {
	return &Command{App: *App}
}

func (c Command) Cmd() error {
	switch arg {
	case "add":
		err := c.CheckArgumentsLength(3)
		if err != nil {
			return e.WrapError(ErrArgumentCount, err)
		}
		id, err := c.App.AddTask(os.Args[2])
		if err != nil {
			return err
		}
		fmt.Printf("Task added successfully (ID: %d)", id)
	case "update":
		err := c.CheckArgumentsLength(4)
		if err != nil {
			return e.WrapError(ErrArgumentCount, err)
		}
		id, err := getId(os.Args[2])
		if err != nil {
			return e.WrapError(fmt.Sprintf(" '%s' is not a number", os.Args[2]), err)
		}
		err = c.App.UpdateTask(id, os.Args[3])
		if err != nil {
			return err
		}
	case "delete":
		err := c.CheckArgumentsLength(3)
		if err != nil {
			return e.WrapError(ErrArgumentCount, err)
		}
		id, err := getId(os.Args[2])
		if err != nil {
			return e.WrapError(fmt.Sprintf(" '%s' is not a number", os.Args[2]), err)
		}
		err = c.App.DeleteTask(id)
		if err != nil {
			return err
		}

	case "list":
		length := len(os.Args)
		if length == 3 {
			switch os.Args[2] {
			case string(models.StatusDone):
				tasks, _ := c.App.ListByStatus(models.StatusDone)
				c.Table(tasks)
			case string(models.StatusInProgress):
				tasks, _ := c.App.ListByStatus(models.StatusInProgress)
				c.Table(tasks)
			case string(models.StatusTodo):
				tasks, _ := c.App.ListByStatus(models.StatusTodo)
				c.Table(tasks)
			default:
				err := e.WrapError(strings.Join(os.Args, " "), errors.New(InvalidCmd))
				if err != nil {
					fmt.Printf("Try:\n-'todo';\n-'in-progress;\n-'done';\n")
				}
				return err

			}
		} else {
			err := c.CheckArgumentsLength(2)
			if err != nil {
				return e.WrapError(ErrArgumentCount, err)
			}
			tasks, _ := c.App.ListAllTasks()
			c.Table(tasks)
		}
	case "mark-in-progress":
		err := c.CheckArgumentsLength(3)
		if err != nil {
			return e.WrapError(ErrArgumentCount, err)
		}
		id, err := getId(os.Args[2])
		if err != nil {
			return e.WrapError(fmt.Sprintf(" '%s' is not a number", os.Args[2]), err)
		}
		err = c.App.MarkTask(id, models.StatusInProgress)
		if err != nil {
			return err
		}
	case "mark-done":
		err := c.CheckArgumentsLength(3)
		if err != nil {
			return e.WrapError(ErrArgumentCount, err)
		}
		id, _ := getId(os.Args[2])
		err = c.App.MarkTask(id, models.StatusDone)
		if err != nil {
			return err
		}
	case "mark-todo":
		err := c.CheckArgumentsLength(3)
		if err != nil {
			return e.WrapError(ErrArgumentCount, err)
		}
		id, err := getId(os.Args[2])
		if err != nil {
			return e.WrapError(fmt.Sprintf(" '%s' is not a number", os.Args[2]), err)
		}
		err = c.App.MarkTask(id, models.StatusTodo)
		if err != nil {
			return err
		}

	default:
		err := e.WrapError(strings.Join(os.Args, " "), errors.New(InvalidCmd))
		if err != nil {
			fmt.Printf("List of commands:\n\033[34madd <«your_task»> \nupdate <task_id> <«your_task»>\ndelete <task_id>\nmark-in-progress <task_id>\nmark-done <task-id>\nmark-todo <task-id>\nlist #will list all tasks\nlist done \nlist todo\nlist in-progress\u001B[0m\n")
		}
		return err
	}
	return nil
}

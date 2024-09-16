package cmd

import (
	"errors"
	"fmt"
	"os"
	"task-tracker/app"
	"task-tracker/lib/color"
	"task-tracker/models"
	"text/tabwriter"
)

type CLI interface {
	Cmd() error
	Table(tasks models.Task) error
	CheckArgumentsLength(int) error
}

type Command struct {
	App app.App
}

func (c Command) Table(tasks []models.Task) error {
	const (
		idWidth      = 5
		descWidth    = 30
		statusWidth  = 15
		createdWidth = 25
		updatedWidth = 25
	)
	if tasks == nil || len(tasks) == 0 {
		fmt.Println(color.Yellow + "No tasks found." + color.Reset)
	} else {
		w := tabwriter.NewWriter(os.Stdout, 10, 10, 2, ' ', tabwriter.Debug)
		fmt.Fprintf(w, fmt.Sprintf("%%-%ds\t%%-%ds\t%%-%ds\t%%-%ds\t%%-%ds\n", idWidth, descWidth, statusWidth, createdWidth, updatedWidth),
			"ID", "DESCRIPTION", "STATUS", "CREATED", "UPDATED")
		for _, task := range tasks {
			switch task.Status {
			case models.StatusDone:
				task.Status = color.Green + models.StatusDone + color.Reset
			case models.StatusTodo:
				task.Status = color.Magenta + models.StatusTodo + color.Reset
			case models.StatusInProgress:
				task.Status = color.Blue + models.StatusInProgress + color.Reset
			}
			fmt.Fprintf(w, fmt.Sprintf("%%-%dd\t%%-%ds\t%%-%ds\t%%-%ds\t%%-%ds\n", idWidth, descWidth, statusWidth, createdWidth, updatedWidth),
				task.Id, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt)
		}
		return w.Flush()
	}
	return nil
}

// CheckArgumentsLength check quantity of args
func (c Command) CheckArgumentsLength(q int) error {
	if q < len(os.Args) {
		return errors.New("too many arguments")
	} else if q > len(os.Args) {
		return errors.New("too few arguments")
	} else {
		return nil
	}
}

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
	Table(listFunc ListFunc) error
	CheckArguments(int) error
}

type Command struct {
	App app.App
}

type ListFunc func() ([]models.Task, error)

func (c Command) Table(listFunc ListFunc) error {
	taskFile, err := listFunc()
	if err != nil {
		return err
	}
	if taskFile == nil {
		fmt.Println(color.Yellow + "No tasks found." + color.Reset)
	} else {
		w := tabwriter.NewWriter(os.Stdout, 0, 20, 1, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "ID\tDESCRIPTION\tSTATUS\tCREATED\tUPDATED")
		for _, task := range taskFile {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
		w.Flush()
	}
	return nil
}

// CheckArguments check quantity of args
func (c Command) CheckArguments(q int) error {
	if q < len(os.Args) {
		return errors.New("too many arguments")
	} else if q > len(os.Args) {
		return errors.New("too few arguments")
	} else {
		return nil
	}
}

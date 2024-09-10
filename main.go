package main

import (
	"fmt"
	"os"
	"task-tracker/app"
	"task-tracker/cmd"
	"task-tracker/lib/color"
	"task-tracker/storage/files"
)

func main() {
	storage := files.NewFileStorage("tasks.json")
	app1 := app.NewApp(storage)
	cmd1 := cmd.NewCmd(app1)
	err := cmd1.Cmd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError: %v%s", color.Red, err, color.Reset)
	}

}

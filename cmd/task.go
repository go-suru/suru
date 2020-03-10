package cmd

import "errors"

type Task struct{}

func (Task) Cmd(_ Context) error {
	return errors.New("Cmder not implemented for Task")
}

func (Task) Short() string { return "Schedule a new Task" }
func (Task) Help() string  { return "TODO" }

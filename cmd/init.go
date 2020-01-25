package cmd

import "errors"

type Init struct{}

func (Init) Cmd() error {
	return errors.New("Cmder not implemented for Init")
}

func (Init) Short() string { return "Initialize Suru for a repo" }
func (Init) Help() string  { return "TODO" }

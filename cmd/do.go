package cmd

import "errors"

type Do struct{}

func (Do) Cmd() error {
	return errors.New("Cmder not implemented for Do")
}

func (Do) Short() string { return "Do some Task" }
func (Do) Long() string  { return "TODO" }

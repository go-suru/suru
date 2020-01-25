package cmd

import "errors"

type Mode struct{}

func (Mode) Cmd() error {
	return errors.New("Cmder not implemented for Mode")
}

func (Mode) Short() string { return "Set the Mode (default Private)" }
func (Mode) Help() string  { return "TODO" }

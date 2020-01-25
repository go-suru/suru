package cmd

import "errors"

type Help struct{}

func (Help) Cmd() error {
	return errors.New("Cmder not implemented for Help")
}

func (Help) Short() string { return "Help for Suru commands" }
func (Help) Help() string  { return "TODO" }

package cmd

import "errors"

type Sub struct{}

func (Sub) Cmd(_ Context) error {
	return errors.New("Cmder not implemented for Sub")
}

func (Sub) Short() string { return "Subscribe to a Suru channel" }
func (Sub) Help() string  { return "TODO" }

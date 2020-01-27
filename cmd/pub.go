package cmd

import "errors"

type Pub struct{}

func (Pub) Cmd(_ Context) error {
	return errors.New("Cmder not implemented for Pub")
}

func (Pub) Short() string { return "Publish to a Suru channel" }
func (Pub) Help() string  { return "TODO" }

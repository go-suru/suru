package cmd

import "errors"

type Config struct{}

var (
	_ = Cmder(Config{})
	_ = Helper(Config{})
)

func (Config) Cmd(_ Context) error {
	return errors.New("Cmder not implemented for Config")
}

func (Config) Short() string { return "Configure Suru" }
func (Config) Help() string  { return "TODO" }

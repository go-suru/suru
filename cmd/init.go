package cmd

import (
	"os"
	"path/filepath"

	"gopkg.in/suru.v0/config"

	"github.com/pkg/errors"
)

type Init struct{}

func (Init) Cmd(ctx Context) error {
	// If initializing without a directive, use local .suru.
	cfg := ctx.Self.Config
	if cfg == "" {
		cfg = ".suru"
	}
	if ctx.Data == "" {
		ctx.Data = ".suru"
	}
	c := filepath.Join(cfg, "config.toml")
	_, err := os.Stat(c)
	switch {
	case os.IsNotExist(err):
		if err := os.MkdirAll(cfg, 0744); err != nil {
			return errors.Wrapf(err, "creating folder for %s", c)
		}
	case err == nil:
	default:
		return errors.Wrapf(err, "checking %s for suru config", c)
	}

	return config.Write(ctx.Config.Merge(config.Defaults()), c)
}

func (Init) Short() string { return "Initialize Suru for a repo" }
func (Init) Help() string  { return "TODO" }

package cmd

import (
	"gopkg.in/suru.v0/cui"
	
	"github.com/pkg/errors"
	"github.com/jroimartin/gocui"
)

type Live struct{}

func (Live) Cmd() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return errors.Wrap(err, "creating live UI")
	}

	return cui.Live(g)
}

func (Live) Short() string { return "Enter live mode" }
func (Live) Help() string  { return "TODO" }

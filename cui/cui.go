package cui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

func Live(g *gocui.Gui) error {
	defer g.Close()

	g.SetManagerFunc(Layout)

	err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit)
	if err != nil {
		return errors.Wrap(err, "setting keybindings")
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return errors.Wrap(err, "main UI loop")
	}

	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView(
		"hello",
		maxX/2-7,
		maxY/2,
		maxX/2+7,
		maxY/2+2,
	); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

package cui

import (
	"gopkg.in/suru.v0/cui/renderer"
	"gopkg.in/suru.v0/cui/state"

	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

type State state.State

func (s *State) Layout(g *gocui.Gui) error {
	ctx := context{
		Gui: g,
		Context: renderer.Context{
			Dim: renderer.DimFromGui(g),
		},
	}

	if err := ctx.withRenderer(s.Root).Render(); err != nil {
		return errors.Wrap(err, "rendering root view")
	}

	if p := s.Popover; p != nil {
		err := ctx.withRenderer(p).withDim(ctx.HalfCtr()).Render()
		if err != nil {
			return errors.Wrap(err, "rendering popover")
		}
	}

	return nil
}

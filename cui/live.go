package cui

import (
	"gopkg.in/suru.v0/cui/state"
	"gopkg.in/suru.v0/cui/view"

	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

func (s *State) Live(g *gocui.Gui) error {
	defer g.Close()

	var ss *state.State = (*state.State)(s)

	g.SetManager(s)

	for _, kb := range []keyBinding{{
		Handler: &view.Help{Root: ss},
		keys: []hotkey{
			{'h', gocui.ModNone},
			{'?', gocui.ModNone},
		},
	}, {
		Handler: HandlerFn(Quit),
		keys: []hotkey{
			{gocui.KeyCtrlC, gocui.ModNone},
			{'q', gocui.ModNone},
		},
	}} {
		if err := kb.set(g); err != nil {
			return errors.Wrap(err, "setting keybindings")
		}
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return errors.Wrap(err, "main UI loop")
	}

	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

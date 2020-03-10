package cui

import (
	"gopkg.in/suru.v0/cui/state"
	"gopkg.in/suru.v0/cui/view"

	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

func (s *State) Live(g *gocui.Gui) error {
	defer g.Close()

	var ss *state.State = &s.State

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

	// TODO: Respect SIGINT even if caught in a dead loop.
	//       Probably just put a Context everywhere and insist on
	//       checking its Done channel.
	// signal.Notify(os.Interrupt)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return errors.Wrap(err, "main UI loop")
	}

	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

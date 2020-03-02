package view

import (
	"gopkg.in/suru.v0/cui/renderer"
	"gopkg.in/suru.v0/cui/state"

	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

// Help is a handler for the help interface.
type Help struct {
	Root *state.State
	View
}

func (h Help) Key() renderer.VKey { return KeyHelp }

// handle implements cui.Handler on Help.  It toggles the help overlay.
func (h *Help) Handle(g *gocui.Gui, v *gocui.View) error {
	root := h.Root

	switch h.State {
	case StateUninitialized:
		// Set up the Frame.
		h.Renderer = &renderer.Frame{
			VKey:     h.Key(),
			Stringer: h,
			Framed:   true,
			Title:    "Help",
		}

		// Attach it to the root State as its Popover.
		root.Popover = h

		// Update the View lifecycle.
		h.State = StateVisible

	case StateVisible:
		// Tear down and detach the Frame.
		if err := h.Teardown(g); err != nil {
			return errors.Wrap(err, "handling teardown")
		}

		root.Popover = nil

		// Update the View lifecycle.
		h.State = StateHidden

	case StateHidden:
		// Un-hide and attach the Frame.
		root.Popover = h

		// Update the View lifecycle.
		h.State = StateVisible
	}

	// No problem!
	return nil
}

func (h Help) String() string {
	return `
Hotkeys:
	(h) Toggle Help
	(q) Quit
	`
}

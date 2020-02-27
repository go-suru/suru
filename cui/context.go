package cui

import (
	"gopkg.in/suru.v0/cui/renderer"

	"github.com/jroimartin/gocui"
)

// context is an ephemeral wrapper for the relevant values for Layout.
// It should be rebuilt in every Layout call.
type context struct {
	*gocui.Gui
	renderer.Context
	renderer.Renderer
}

// with returns a copy of the context with the given View.
func (c context) withRenderer(r renderer.Renderer) context {
	c.Renderer = r
	return c
}

func (c context) withDim(d renderer.Dim) context {
	c.Context.Dim = d
	return c
}

// Render calls the current Renderer in the context.
func (c context) Render() error {
	return c.Renderer.Render(c.Context, c.Gui)
}

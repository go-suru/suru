//package Renderer implements gocui helpers and view component types.
package renderer

import (
	"github.com/jroimartin/gocui"
)

type Renderer interface {
	Render(Context, *gocui.Gui) error
	Teardown(*gocui.Gui) error
}

type Context struct {
	Dim
}

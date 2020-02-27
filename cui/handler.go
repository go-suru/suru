package cui

import "github.com/jroimartin/gocui"

type Handler interface {
	Handle(*gocui.Gui, *gocui.View) error
}

type HandlerFn func(*gocui.Gui, *gocui.View) error

func (f HandlerFn) Handle(g *gocui.Gui, v *gocui.View) error {
	return f(g, v)
}

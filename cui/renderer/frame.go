package renderer

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

type Frame struct {
	*gocui.View
	VKey

	fmt.Stringer

	Hidden bool
	Framed bool
	Title  string
}

func (f *Frame) Render(c Context, g *gocui.Gui) (err error) {
	do, ds := c.Origin, c.Origin.Add(c.Size)
	v, hid, ks := f.View, f.Hidden, f.VKey.String()

	ox, oy, sx, sy := do.X, do.Y, ds.X-1, ds.Y-1

	switch {
	case v != nil && hid:
		// The user wants to hide / destroy the view.
		return f.Teardown(g)

	case v == nil && !hid:
		// The view hasn't been initialized yet, so create and
		// cache it using the gocui API.
		v, err = g.SetView(ks, ox, oy, sx, sy)
		if err != nil && err != gocui.ErrUnknownView {
			// ErrUnknownView is normal here for new views.
			return errors.Wrapf(err, "creating view %s", ks)
		}

		v.Frame = f.Framed
		v.Title = f.Title

		f.View = v

		// TODO: handle case where frame view was set but deleted.

	case v != nil && !hid:
		// The view is already initialized and is visible.
		if v, err = g.SetView(ks, ox, oy, sx, sy); err != nil {
			return errors.Wrapf(err, "setting current view %s", ks)
		}

	default:
		// Not visible, not created, or something.  No problem.
		return nil
	}

	content := Text{W: sx - ox, H: sy - oy, Stringer: f.Stringer}
	v.Clear()
	content.WriteTo(v)
	return nil
}

func (f *Frame) Teardown(g *gocui.Gui) error {
	ks := f.VKey.String()
	f.View = nil
	return errors.Wrapf(g.DeleteView(ks), "tearing down %s", ks)
}

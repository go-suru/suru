package cui

import (
	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

type hotkey struct {
	// k must be one of rune or gocui.Key.
	k interface{}
	gocui.Modifier
}

type keyBinding struct {
	Handler
	view string
	keys []hotkey
}

func (k keyBinding) set(g *gocui.Gui) error {
	for _, kk := range k.keys {
		switch kk.k.(type) {
		case gocui.Key:
		case rune:
		default:
			return errors.Errorf("key type must be rune "+
				"or gocui.Key, was %T", kk.k,
			)
		}

		err := g.SetKeybinding(
			k.view,
			kk.k,
			kk.Modifier,
			k.Handle,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

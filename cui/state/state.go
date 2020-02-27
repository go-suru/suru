package state

import "gopkg.in/suru.v0/cui/renderer"

// TODO: State serialization / loader to let user jump right back in.

// State models the state of the UI.
type State struct {
	Root renderer.Renderer
	renderer.Popover
}

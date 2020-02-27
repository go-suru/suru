package view

import "gopkg.in/suru.v0/cui/renderer"

// View is a renderer for gocui Views.
type View struct {
	State
	renderer.Renderer
}

type State int

const (
	StateUninitialized State = iota
	StateVisible
	StateHidden
)

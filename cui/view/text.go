package view

import (
	"fmt"
)

// Text is a Contenter for rendering a string.
type Text string

func Sprintf(f string, args ...interface{}) Text {
	return Text(fmt.Sprintf(f, args...))
}

func (t Text) String() string { return string(t) }

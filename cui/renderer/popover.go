package renderer

import "fmt"

type Popover interface {
	Renderer
	fmt.Stringer
}

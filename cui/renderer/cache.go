package renderer

import (
	"fmt"
)

type VKey int

func (v VKey) String() string { return fmt.Sprintf("%d", v) }

type Keyer interface{ Key() VKey }

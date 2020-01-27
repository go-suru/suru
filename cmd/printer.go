package cmd

import (
	"fmt"

	"github.com/pkg/errors"
)

type Printer string

func (p Printer) Cmd(ctx Context) (err error) {
	_, err = fmt.Fprintln(ctx, p)
	return errors.Wrap(err, "printing")
}

package cmd

import (
	"bufio"

	"github.com/pkg/errors"
)

type ParseErr struct{ error }

func IsParseErr(err error) (is bool) {
	_, is = err.(ParseErr)
	return
}

// Helper defines a short description and a longer help message for use
// in a help page.
type Helper interface {
	Short() string
	Help() string
}

// Parser consumes command-line arguments to return a Cmder.
// Idea: Use Parser with Help to validate example usage!
type Parser interface {
	Parse(arg string) (Cmder, error)
}

type Context struct {
	*bufio.Writer
}

// Cmder is a simple CLI command interface.
type Cmder interface {
	Cmd(Context) error
}

// CmdFn is a function implementation of Cmder.
type CmdFn func(Context) error

// Cmd implements Cmder on CmdFn.
func (c CmdFn) Cmd(cx Context) error { return c(cx) }

var help = new(Help)
var cmds = map[string]Cmder{
	// User wants to enter interactive mode.
	"live": Live{},
	// User wants to add Suru metadata and hooks to a repo.
	"init": Init{},
	// User wants to configure Suru.
	"config": Config{},
	// User wants to decide what mode (public, etc.) to use.
	"mode": Mode{},
	// User wants to start on a task.
	"do": Do{},
	// User wants to define or inspect a task.
	"task": Task{},
	// User wants to know the Suru version number.
	"version": Version{},
	// User wants to publish his work.
	"pub": Pub{},
	// User wants to subscribe to a Suru stream.
	"sub": Sub{},
	// User wants help.
	"help": help, "h": help, "?": help,
}

// Parse consumes a list of command-line arguments to return a Cmder.
func Parse(args ...string) (cc Cmder, err error) {
	c, ok := cmds[args[0]]
	if !ok {
		err = ParseErr{
			errors.Errorf("unknown cmd %#q", args[0]),
		}
		return
	}
	for i := 1; i < len(args); i++ {
		if p, ok := c.(Parser); !ok {
			break
		} else if c, err = p.Parse(args[i]); err != nil {
			return nil, errors.Wrapf(err, "arg %d", i)
		}
	}
	return c, nil
}

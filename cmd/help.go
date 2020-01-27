package cmd

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/pkg/errors"
)

type Help struct{ Helper }

var (
	h = new(Help)

	_ = Cmder(h)
	_ = Helper(h)
	_ = Parser(h)

	_ = sort.Interface(ByLength(nil))
)

func (h *Help) Cmd(ctx Context) (err error) {
	hh := h.Helper
	if hh == nil {
		// TODO: Iterate and append inner helpers
		hh = h
	}

	_, err = fmt.Fprintf(ctx, "%s\n\n%s\n",
		hh.Short(), hh.Help(),
	)
	return
}

func (h *Help) Parse(next string) (cc Cmder, err error) {
	if cm, ok := cmds[next]; !ok {
		err = ParseErr{
			errors.Errorf("unknown topic %#q", next),
		}
		return
	} else if hh, ok := cm.(Helper); !ok {
		err = ParseErr{
			errors.Errorf("topic %#q has no help", next),
		}
		return
	} else {
		h.Helper = hh
	}

	return h, nil
}

func (Help) Short() string {
	return "Help for Suru commands (with shortcuts)"
}

type ByLength []string

func (b ByLength) Len() int      { return len(b) }
func (b ByLength) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b ByLength) Less(i, j int) bool {
	bi, bj := b[i], b[j]
	il, jl := len(bi), len(bj)
	return il < jl || (il == jl && b[i] < b[j])
}

func (Help) Help() (is string) {
	var (
		helps []string
		names = make(map[Cmder][]string)

		buf bytes.Buffer
		tw  = tabwriter.NewWriter(
			&buf, 0, 8, 0, '\t', tabwriter.TabIndent,
		)
	)

	// Collect all the duplicated names of commands.
	// For example, "Help" is named "help", "h", and "?".
	// The first name in each slice is the longest for that Cmder.
	for name, c := range cmds {
		names[c] = append(names[c], name)
		sort.Sort(sort.Reverse(ByLength(names[c])))
	}

	for _, aliases := range names {
		helps = append(helps, aliases[0])
	}

	sort.Strings(helps)

	for _, s := range helps {
		cc := cmds[s]
		name, rest := names[cc][0], names[cc][1:]
		desc := "Not documented"

		if h, ok := cmds[s].(Helper); ok {
			desc = h.Short()
		}

		text := fmt.Sprintf(" - %s%s\t%s\n",
			name,
			func() string {
				if len(rest) == 0 {
					return ""
				}
				return fmt.Sprintf(" (%s)",
					strings.Join(rest, ", "),
				)
			}(),
			desc,
		)
		_, _ = tw.Write([]byte(text))
	}

	tw.Flush()

	return buf.String()
}

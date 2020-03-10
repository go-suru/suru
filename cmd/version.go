package cmd

import "gopkg.in/suru.v0"

type Version struct{}

var (
	v Version

	_ = Helper(v)
	_ = Cmder(v)
)

func (v Version) Cmd(c Context) error {
	return Printer("Suru v" + suru.Version).Cmd(c)
}

func (Version) Short() string {
	return "Print the Suru version (in SemVer 2.0)"
}

func (Version) Help() (is string) {
	return "Print the Suru version (in SemVer 2.0)"
}

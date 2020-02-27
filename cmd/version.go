package cmd

import "gopkg.in/suru.v0"

type Version struct{}

var (
	v Version

	_ = Helper(v)
	_ = Parser(v)
)

func (v Version) Cmd(_ Context) error { return nil }


func (v Version) Parse(next string) (cc Cmder, err error) {
	return Printer("Suru v"+suru.Version), nil
}

func (Version) Short() string {
	return "Print the Suru version (in SemVer 2.0)"
}

func (Version) Help() (is string) {
	return "Print the Suru version (in SemVer 2.0)"
}

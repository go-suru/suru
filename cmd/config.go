package cmd

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"gopkg.in/suru.v0/config"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func validFields(t interface{}) map[string]bool {
	m := make(map[string]bool)
	for _, f := range fields(t) {
		m[f] = true
	}
	return m
}

func fields(t interface{}) (is []string) {
	tc := reflect.TypeOf(t)
	n := tc.NumField()
	for i := 0; i < n; i++ {
		is = append(is, strings.ToLower(tc.Field(i).Name))
	}

	sort.Strings(is)
	return
}

func index(t interface{}, by ...string) interface{} {
	tv := reflect.ValueOf(t)
	for _, f := range by {
		tv = tv.FieldByName(f)
	}
	return tv.Interface()
}

type Config struct{}

type ConfigGet []string
type ConfigSet struct {
	key   []string
	value string
}

type ConfigLister struct{}

var (
	_ = Parser(new(Config))
	_ = Cmder(Config{})

	_ = Parser(new(ConfigGet))
	_ = Cmder(new(ConfigGet))

	_ = Parser(new(ConfigSet))
	_ = Cmder(new(ConfigSet))

	_ = Helper(Config{})
)

func (c *Config) Parse(arg string) (Cmder, error) {
	switch arg {
	case "get":
		return new(ConfigGet), nil
	case "set":
		// return new(ConfigSet), nil
	case "path":
		return &ConfigGet{"Self", "Config"}, nil
	}

	return nil, errors.Errorf("No such Config command %q", arg)
}

func (c *ConfigGet) Parse(arg string) (is Cmder, err error) {
	if strings.Contains(arg, ".") {
		is = c
		var p Parser = c
		for _, f := range strings.Split(arg, ".") {
			if is, err = p.Parse(f); err != nil {
				return Printer(fmt.Sprintf(
					"Invalid path: %s", err,
				)), nil
			}
			p = is.(Parser)
		}
		return
	}

	if !validFields(index(config.Config{}, *c...))[arg] {
		return nil, errors.Errorf("invalid field %q", arg)
	}

	tmp := ConfigGet(append(*c,
		strings.ToUpper(string(arg[0]))+arg[1:],
	))
	return &tmp, nil
}

func (c *ConfigGet) Cmd(ct Context) error {
	fv := reflect.ValueOf(ct.Config)
	for _, ff := range *c {
		fv = fv.FieldByName(ff)
	}

	bs, err := yaml.Marshal(fv.Interface())
	if err != nil {
		return errors.Wrapf(err, "marshaling YAML")
	}

	return Printer(bs).Cmd(ct)
}

func (c *ConfigSet) Parse(next string) (Cmder, error) {
	return nil, errors.New("Parse not implemented")
}

func (c *ConfigSet) Cmd(ct Context) error {
	return Printer("Choose a parameter:" +
		strings.Join(fields(index(
			config.Config{}, (*c).key...,
		)), "\n  - "),
	).Cmd(ct)
}

func (c Config) Cmd(ct Context) error {
	return Printer(c.Help()).Cmd(ct)
}

func (Config) Short() string { return "Configure Suru" }
func (Config) Help() string {
	return `
Config subcommands

 - get [key, key...]	Get the config value at the given path.
 - set [k1.k2 value]	Set the config value at the given path (TODO.)
 - path			Print the path to the current config.
`[1:]
}

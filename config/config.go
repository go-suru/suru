package config

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Self struct {
	Data   string `toml:"data"`
	Config string `toml:"-"`
}

func (s Self) Defaults() Self {
	return Self{Data: DefaultDataPath()}
}

func (s Self) Merge(from Self) (is Self) {
	is = s
	if len(is.Data) == 0 {
		is.Data = from.Data
	}
	if len(is.Config) == 0 {
		is.Config = from.Config
	}

	return
}

type Pub struct {
	Host string `toml:"host"`
}

func (p Pub) Defaults() Pub {
	return Pub{Host: "https://go-suru.github.io"}
}

func (p Pub) Merge(from Pub) (is Pub) {
	is = p
	if len(is.Host) == 0 {
		is.Host = from.Host
	}

	return
}

type Topics struct{}

func (t Topics) Defaults() Topics {
	return Topics{}
}

func (t Topics) Merge(from Topics) (is Topics) {
	return is
}

type Config struct {
	Self   `toml:"self"`
	Pub    `toml:"pub,omitempty"`
	Topics `toml:"topics,omitempty"`
}

func Defaults() Config {
	return Config{
		Self:   Self{}.Defaults(),
		Pub:    Pub{}.Defaults(),
		Topics: Topics{}.Defaults(),
	}
}

// DefaultConfigPaths returns the default path to search for
// config.toml: $HOME/.config/suru (or %APPDATA%/suru on Windows.)
func DefaultConfigPath() string {
	home, err := os.UserConfigDir()
	if err != nil {
		return ".suru"
	}

	// /home/user/.config/suru
	return filepath.Join(home, "suru")
}

func DefaultDataPath() (is string) {
	if runtime.GOOS == "windows" {
		return filepath.Join(
			os.Getenv("APPDATA"),
			"suru",
		)
	}

	if dh := os.Getenv("XDG_DATA_HOME"); len(dh) > 0 {
		return filepath.Join(is, "suru")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return filepath.Join(".suru")
	}

	return filepath.Join(home, "suru")
}

// Load searches the given directories for config.toml and loads it into
// the given Config.  If there is no config.toml found in any of the
// paths, it creates it in the first location.  If no location is
// provided, it uses ".suru".
func Load(cfg *Config, hints ...string) (err error) {
	if len(hints) == 0 {
		hints = []string{".suru", DefaultConfigPath()}
	}

	var inf os.FileInfo
	var p string
findFolder:
	for _, p = range hints {
		inf, err = os.Stat(p)
		switch {
		case os.IsNotExist(err):
			continue
		case err == nil && !inf.IsDir():
			// That wasn't a folder!
			return errors.Wrapf(err, "%s isn't a directory", p)
		case err == nil:
			// Found a folder, check here for config file.
			_, err = os.Stat(filepath.Join(p, "config.toml"))
			switch {
			case os.IsNotExist(errors.Cause(err)):
				// No config here, try another folder.
				continue
			case err == nil:
				break findFolder
			default:
				return errors.Wrapf(err, "checking %s", p)
			}
		default:
			return errors.Wrapf(err, "opening %s", p)
		}
	}

	if os.IsNotExist(err) {
		// The suru config wasn't found anywhere, so use the default
		// config path.
		if cfgPath := cfg.Config; len(cfgPath) > 0 {
			// User wants it in a particular place.
			p = cfgPath
		} else {
			p = DefaultConfigPath()
		}
	}

	var f *os.File
	p = filepath.Join(p, "config.toml")

	// Try opening the config file.
	f, err = os.Open(p)
	switch {
	case os.IsNotExist(err):
		// Create and initialize the config file.
		return Write(cfg.Merge(Defaults()), p)
	case err == nil:
	default:
		return errors.Wrapf(err, "opening %s", p)
	}

	// If found, load and merge it.
	var newCfg Config
	if err := LoadFrom(f, &newCfg); err != nil {
		return errors.Wrap(err, "loading config from file")
	}

	// Let the user know the path, even though it's omitted.
	*cfg = cfg.Merge(newCfg)
	cfg.Config = p
	return nil
}

func (c Config) Merge(from Config) (is Config) {
	is = c
	is.Self = is.Self.Merge(from.Self)
	is.Pub = is.Pub.Merge(from.Pub)
	is.Topics = is.Topics.Merge(from.Topics)

	return
}

func LoadFrom(r io.Reader, into *Config) error {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "reading")
	}

	return errors.Wrap(toml.Unmarshal(bs, into), "decoding TOML")
}

func Write(cfg Config, path string) (err error) {
	_, err = os.Stat(filepath.Dir(path))
	switch {
	case os.IsNotExist(err):
		if err = os.MkdirAll(path, 0744); err != nil {
			return errors.Wrapf(err, "making %s", path)
		}
	case err == nil:
	default:
		return errors.Wrapf(err, "checking %s", path)
	}

	var f *os.File
	if f, err = os.Create(path); err != nil {
		return errors.Wrapf(err, "creating %s", path)
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(cfg)
	return errors.Wrapf(err, "encoding config to %s", path)
}

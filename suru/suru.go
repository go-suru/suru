package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/suru.v0"
	"gopkg.in/suru.v0/cmd"
	"gopkg.in/suru.v0/config"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

var (
	cfgPath, dbPath string
)

func initVars() {
	flag.StringVar(&cfgPath, "cfg-path",
		"",
		"The Suru config folder",
	)
	flag.StringVar(&dbPath, "db-path",
		"",
		"The Suru database object",
	)

	flag.Parse()
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w,
		"Suru v%s\n\n"+
			"Usage:\n%s",
		suru.Version, cmd.Help{}.Help(),
	)
}

func main() {
	initVars()

	if len(os.Args) < 2 {
		printUsage(os.Stderr)
		os.Exit(1)
	}

	cm, err := cmd.Parse(flag.Args()...)
	switch {
	case cmd.IsParseErr(errors.Cause(err)):
		fmt.Fprintf(os.Stderr, "Parsing failed: %s\n\n", err)
		printUsage(os.Stderr)
		os.Exit(1)
	case err != nil:
		log.Fatalf("Parsing failed: %s\n", err)
		printUsage(os.Stderr)
		os.Exit(1)
	}

	cfg := config.Config{
		Self: config.Self{
			Data:   dbPath,
			Config: cfgPath,
		},
	}

	// Special-case for Init: don't load config.
	if _, ok := cm.(cmd.Init); !ok {
		var paths []string
		if cfgPath != "" {
			paths = []string{cfgPath}
		}
		if err := config.Load(&cfg, paths...); err != nil {
			log.Fatalf("Loading config failed: %s", err)
		}
	}

	var db *bolt.DB
	if cdb, ok := cm.(cmd.DB); ok {
		_ = cdb
		path := filepath.Join(cfg.Data, "suru.db")
		db, err := bolt.Open(path, 0600, new(bolt.Options))
		if err != nil {
			log.Fatalf("Opening DB failed: %s", err)
		}
		defer db.Close()
	}

	ctx := cmd.Context{
		Writer: bufio.NewWriter(os.Stderr),
		DB:     db,
		Config: cfg,
	}
	if err := cm.Cmd(ctx); err != nil {
		log.Fatalf("Command failed: %s", err)
	}

	ctx.Flush()
	// Parse CLI args
	// Load config
	// TODO: Encrypt DB at rest.
	// Set up peer connection & event source (if any)
	// Updates available? (if online mode)
	//   - How to make this non-intrusive and optional
	//     but still easy?
	// If "live" mode, load UI

	// Options:
	// Create or load Topic (local dir)
	// Create or connect to Stream
	// Create Workshop Item
	// Schedule Workshop Item

	// Connect
	//  - Generate ID keypair
	//  - Private mode
	//  - Org mode
	//  - Public mode

	// Collab
	//  - Chat
	//  - Profile
	//  - Workshop
	//  - Issues
	//  - Schedule
	//  - Worklog

	// Decentralize
	//  - Tracker (IP)
	//  - Tracker (Tor)
	//  - Tracker (Peer)
	//  - Locate peer
	//  - Connect to peer
	//  - Update peer
	//  - Forward peer update

}

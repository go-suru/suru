package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/suru.v0"
	"gopkg.in/suru.v0/cmd"

	"github.com/pkg/errors"
)

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage:\n%s", cmd.Help{}.Help())
}

func main() {
	fmt.Fprintf(os.Stderr, "Suru v%s\n\n", suru.Version)

	if len(os.Args) < 2 {
		printUsage(os.Stderr)
		os.Exit(1)
	}

	cm, err := cmd.Parse(os.Args[1:]...)
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

	ctx := cmd.Context{Writer: bufio.NewWriter(os.Stderr)}
	if err := cm.Cmd(ctx); err != nil {
		log.Fatalf("Command failed: %s", err)
	}

	ctx.Flush()
	// Parse CLI args
	// Load config
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

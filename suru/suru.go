package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/suru.v0/cmd"
)

const help = `suru v0.0.0`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s\n\nInsufficient args", help)
		os.Exit(1)
	}

	if cm, err := cmd.Parse(os.Args[1:]...); err != nil {
		log.Fatalf("Failed to parse args: %s\n\n%s", err, help)
	} else if err := cm.Cmd(); err != nil {
		log.Fatalf("Command failed: %s", err)
	}

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

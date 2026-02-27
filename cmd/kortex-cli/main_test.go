package main

import (
	"os"
	"testing"
)

func TestMain_VersionSubcommand(t *testing.T) {
	// Save original os.Args and restore after test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Set os.Args to call the version subcommand
	os.Args = []string{"kortex-cli", "version"}

	// Call main() - test passes if it doesn't panic
	main()
}

package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestRootCmd_HasVersionCommand(t *testing.T) {
	versionCmd := rootCmd.Commands()
	found := false
	for _, cmd := range versionCmd {
		if cmd.Name() == "version" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected rootCmd to have 'version' subcommand")
	}
}

func TestExecute_WithVersion(t *testing.T) {
	// Save original os.Args and restore after test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Set os.Args to call version subcommand
	os.Args = []string{"kortex-cli", "version"}

	// Redirect output to avoid cluttering test output
	oldStdout := rootCmd.OutOrStdout()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	defer rootCmd.SetOut(oldStdout)

	// Call Execute() - test passes if it doesn't panic
	Execute()
}

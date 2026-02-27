package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestRootCmd_Initialization(t *testing.T) {
	if rootCmd.Use != "kortex-cli" {
		t.Errorf("Expected Use to be 'kortex-cli', got '%s'", rootCmd.Use)
	}

	if rootCmd.Short == "" {
		t.Error("Expected Short description to be set")
	}
}

func TestExecute_WithHelp(t *testing.T) {
	// Save original os.Args and restore after test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Set os.Args to call help
	os.Args = []string{"kortex-cli", "--help"}

	// Redirect output to avoid cluttering test output
	oldStdout := rootCmd.OutOrStdout()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	defer rootCmd.SetOut(oldStdout)

	// Call Execute() - test passes if it doesn't panic
	Execute()
}

func TestExecute_NoArgs(t *testing.T) {
	// Save original os.Args and restore after test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Set os.Args with no subcommand
	os.Args = []string{"kortex-cli"}

	// Redirect output to avoid cluttering test output
	oldStdout := rootCmd.OutOrStdout()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	defer rootCmd.SetOut(oldStdout)

	// Call Execute() - test passes if it doesn't panic
	Execute()
}

# AGENTS.md

This file provides guidance to AI agents when working with code in this repository.

## Project Overview

kortex-cli is a command-line interface for launching and managing AI agents (Claude Code, Goose, Cursor) with custom configurations. It provides a unified way to start different agents with specific settings including skills, MCP server connections, and LLM integrations.

## Build and Test Commands

### Build
```bash
go build ./cmd/kortex-cli
```

### Execute
After building, the `kortex-cli` binary will be created in the current directory:

```bash
# Display help and available commands
./kortex-cli --help

# Execute a specific command
./kortex-cli <command> [flags]
```

### Run Tests
```bash
# Run all tests
go test ./...

# Run tests in a specific package
go test ./pkg/cmd

# Run a specific test
go test -run TestName ./pkg/cmd

# Check code coverage
go test -cover ./...
```

## Architecture

### Command Structure (Cobra-based)
- Entry point: `cmd/kortex-cli/main.go` → calls `pkg/cmd.Execute()`
- Root command: `pkg/cmd/root.go` defines the `kortex-cli` command
- Subcommands: Each command is in `pkg/cmd/<command>.go` and registers itself via `init()`
- Commands use Cobra's pattern: define a `cobra.Command`, register it with `rootCmd.AddCommand()` in `init()`

### Skills System
Skills are reusable capabilities that can be discovered and executed by AI agents:
- **Location**: `skills/<skill-name>/SKILL.md`
- **Claude support**: Skills are symlinked in `.claude/skills/` for Claude Code
- **Format**: Each SKILL.md contains:
  - YAML frontmatter with `name`, `description`, `argument-hint`
  - Detailed instructions for execution
  - Usage examples

### Adding a New Skill
1. Create directory: `skills/<skill-name>/`
2. Create SKILL.md with frontmatter and instructions
3. Symlink in `.claude/skills/`: `ln -s ../../skills/<skill-name> .claude/skills/<skill-name>`

### Adding a New Command
1. Create `pkg/cmd/<command>.go`
2. Define a `cobra.Command` variable
3. Register with `rootCmd.AddCommand()` in the `init()` function
4. Create corresponding test file `pkg/cmd/<command>_test.go`

## Copyright Headers

All source files must include Apache License 2.0 copyright headers with Red Hat copyright. Use the `/copyright-headers` skill to add or update headers automatically. The current year is 2026.

## Dependencies

- Cobra (github.com/spf13/cobra): CLI framework
- Go 1.25.7

## Testing

Tests follow Go conventions with `*_test.go` files alongside source files. Tests use the standard `testing` package and should cover command initialization, execution, and error cases.

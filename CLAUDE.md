# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is `wu`, a Go library containing utility packages for common development tasks. It's organized as a collection of modules rather than a monolithic application.

## Build, Test, and Development Commands

```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./crypto/aes
go test ./moment
go test ./noti

# Run with verbose output
go test -v ./...

# Update dependencies
go mod tidy

# Build the module
go build ./...
```

## Architecture and Package Structure

The library is organized by functional domain into the following key packages:

### Core Runtime (`wu/`)
- `runtime.go` - Runtime utilities (CPU count, goroutines, garbage collection)
- `exit.go` - Exit helpers with messages
- `wu.go` - Global utilities including `TimeBomb()` for app expiration
- `log.go` - Colorful logging with `Log`, `LogErr`, `LogInfo`, `LogWarn`, `LogSuccess`, etc.
- `error.go` - Error handling utilities (`Err`, `OK`, `ErrFatal`, `Errf`)

### CLI Framework (`app/cli/`)
A minimal CLI framework with:
- `app.go` - `App` struct with `Run()`, `End()`, `EndOnError()` entry points
- `command.go` - `Command` struct for subcommands
- `flag.go` - Flag types: `BoolFlag`, `StringFlag`, `IntFlag`
- `context.go` - `Context` for accessing parsed flags and args
- `help.go` - Help templates and `helpCommand`, `versionCommand`

### Utility Packages
- `config/` - JSON config loading (`InitJSON`, `LoadJSON`)
- `crypto/` - MD5, SHA256, AES encryption
- `datetime/` - Date/time utilities
- `encoding/` - Base64, zlib+AES+base64 encoding
- `getpasswd/` - Password input from terminal
- `moment/` - Moment.js-like date library (supports strftime formats)
- `net/` - Network utilities, HTTP client wrapper
- `noti/` - macOS notifications via osascript
- `os/` - File, path, shell, user utilities
- `regex/` - Regex helpers
- `say/` - Text-to-speech (macOS via `say`, Windows via PowerShell)
- `str/` - String conversion utilities
- `structEx/` - Struct reflection helpers (`GetFieldNames`, `GetTagsByArrayPos`, `Explicit`)
- `timezone/` - Timezone-specific time setters
- `wtime/` - Time utilities
- `zip/` - Gzip, zlib, lzma, p7zip compression
- `cron/` - Cron scheduling

### Special Modules
- `termcolor/` - Terminal color output
- `termimg/` - Terminal image display
- `timeout/` - Timeout handling

## Key Patterns

### Logging
The library uses a custom `WuLog` type with colorful terminal output. Call depth is set to `4` by default to skip wrapper functions.

### Error Handling
Uses a global `WuErr` instance with an enable/disable toggle. Functions like `Err()`, `OK()`, `ErrFatal()` return bools indicating error status.

### Time Bomb
`TimeBomb()` in `wu.go` checks against a hardcoded date (default 2026-01-01) and exits if exceeded - used for app expiration.
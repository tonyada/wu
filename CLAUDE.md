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

Behavioral guidelines to reduce common LLM coding mistakes. Merge with project-specific instructions as needed.

**Tradeoff:** These guidelines bias toward caution over speed. For trivial tasks, use judgment.

## 1. Think Before Coding

**Don't assume. Don't hide confusion. Surface tradeoffs.**

Before implementing:
- State your assumptions explicitly. If uncertain, ask.
- If multiple interpretations exist, present them - don't pick silently.
- If a simpler approach exists, say so. Push back when warranted.
- If something is unclear, stop. Name what's confusing. Ask.

## 2. Simplicity First

**Minimum code that solves the problem. Nothing speculative.**

- No features beyond what was asked.
- No abstractions for single-use code.
- No "flexibility" or "configurability" that wasn't requested.
- No error handling for impossible scenarios.
- If you write 200 lines and it could be 50, rewrite it.

Ask yourself: "Would a senior engineer say this is overcomplicated?" If yes, simplify.

## 3. Surgical Changes

**Touch only what you must. Clean up only your own mess.**

When editing existing code:
- Don't "improve" adjacent code, comments, or formatting.
- Don't refactor things that aren't broken.
- Match existing style, even if you'd do it differently.
- If you notice unrelated dead code, mention it - don't delete it.

When your changes create orphans:
- Remove imports/variables/functions that YOUR changes made unused.
- Don't remove pre-existing dead code unless asked.

The test: Every changed line should trace directly to the user's request.

## 4. Goal-Driven Execution

**Define success criteria. Loop until verified.**

Transform tasks into verifiable goals:
- "Add validation" → "Write tests for invalid inputs, then make them pass"
- "Fix the bug" → "Write a test that reproduces it, then make it pass"
- "Refactor X" → "Ensure tests pass before and after"

For multi-step tasks, state a brief plan:
```
1. [Step] → verify: [check]
2. [Step] → verify: [check]
3. [Step] → verify: [check]
```

Strong success criteria let you loop independently. Weak criteria ("make it work") require constant clarification.

---

**These guidelines are working if:** fewer unnecessary changes in diffs, fewer rewrites due to overcomplication, and clarifying questions come before implementation rather than after mistakes.

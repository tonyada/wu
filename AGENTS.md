# AGENTS.md

## Quick Start

```bash
# Build entire library
go build ./...

# Test specific package (e.g., AES encryption)
go test ./crypto/aes

# Run with verbose output
go test -v ./...

# Update dependencies
go mod tidy
```

## Architecture Overview

**Flat single-file package structure.** Each functional domain is a directory containing exactly one `.go` file (except subdirectories for nested packages):

```
wu/              # Root: core runtime utils only (runtime.go, exit.go, wu.go, log.go, error.go)
app/cli/         # Minimal CLI framework (app.go, command.go, flag.go, context.go)
app/cli/elseme/  # TODO code, ignore
<wu.go>          # Legacy file in root (commands/command.go is current entry point)
config/          # JSON config loading
crypto/aes/      # AES implementation (aes.go + aes_test.go)
moment/          # Moment.js-like date lib (moment.go, strftime_parser.go, diff.go)
net/httpclient/  # HTTP client wrapper
noti/            # macOS notifications
os/file/         # File operations
str/             # String utilities
timezone/        # TZ-specific global time setters (beijing/, vancouver/, newyork/)
zip/gzip/        # Gzip compression
```

**Entry points**:
- CLI: `wu/app/cli/app.go` - `App.Run()` / `App.End()` / `App.EndOnError()`
- Runtime: `wu/wu.go` - `TimeBomb()`, logging, error handling

## Critical Constraints

### Time Bomb (DEADLINE)
`wu.TimeBomb()` contains hardcoded expiration date **2026-01-01**. Any code using this will exit on this date. Do not use in production without replacing this date.

### Error Handling Pattern
Error values use `bool` return pattern, not panic:
- `wu.Err(...)`, `wu.OK()`, `wu.ErrFatal(...)` return `bool` indicating error state
- Disable with `wu.WuErr.Disable()` in tests

### Logging
`wu.Log/Err/Info` use colored output via `github.com/muesli/termenv`. Tests disable colors via `os.Getenv("TERM")="dumb`.

## Testing Strategy

**Convention**: `pkg/<name>/<name>_test.go` (parallel to implementation)
- Test files are standalone packages with `package <pkgName>` (not `package testing`)
- Import via `github.com/stretchr/testify/assert`
- No main() entry point; use `_test.go` package convention

**Example**: `/crypto/aes/aes.go` → `/crypto/aes/aes_test.go`

## Package Boundaries

**CLI Commands**: `app/cli/versionCommand`, `app/cli/helpCommand`
**Date/time**: `moment/` (API), `wtime/` (UTC local), `timezone/` (global TZ setter)
**Encoding**: `encoding/base64`, `encoding/zlib_aes_base64` (chained)

## Navigation Shortcuts

- Find all test files: `find . -name '*_test.go'`
- Root entry points: `wu/*.go`, `app/cli/*.go`
- Ignore: `app/cli/elseme/` (legacy todo code)
- Dependencies: `go.mod` uses Go 1.23+ (toolchain 1.24)

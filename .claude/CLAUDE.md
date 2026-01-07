# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is `sz-sdk-go-mock`, a Go package providing mock implementations of Senzing SDK interfaces. It simulates Senzing library functions without requiring actual Senzing C libraries, useful for testing and development. Implements interfaces from `github.com/senzing-garage/sz-sdk-go`.

**Status:** Work-in-progress (0.n.x semantic versions)

## Common Commands

```bash
# Build
make clean build

# Run all tests (sequential execution)
make test

# Run single test
go test -v -run TestSzengine_AddRecord ./szengine

# Run tests with coverage (opens HTML report)
make coverage

# Check coverage meets 75% threshold
make check-coverage

# Lint
make lint

# Auto-fix formatting issues
make fix

# Install dev tools (one-time)
make dependencies-for-development

# Update Go dependencies
make dependencies
```

## Architecture

### Package Structure

| Package             | Purpose                               | Component ID |
|---------------------|---------------------------------------|--------------|
| `szabstractfactory` | Factory creating all mock Sz* objects | 6030         |
| `szconfig`          | In-memory configuration management    | 6031         |
| `szconfigmanager`   | Configuration persistence             | 6032         |
| `szengine`          | Entity resolution & record processing | 6033         |
| `szdiagnostic`      | Repository diagnostics                | 6034         |
| `szproduct`         | License & version info                | 6036         |
| `helper`            | Shared logging/messaging utilities    | -            |
| `testdata`          | Test data helpers                     | -            |

### Mock Object Pattern

Each mock type has configurable `*Result` fields that control return values:

```go
type Szengine struct {
    AddRecordResult           string
    DeleteRecordResult        string
    GetEntityByEntityIDResult string
    // ...
}
```

### Message ID Format

`SZSDKcccceeee` where `cccc` = component ID, `eeee` = message ID (e.g., `SZSDK60330001`)

### Standard Patterns in Implementation

1. **Trace logging** with `traceEntry`/`traceExit`
2. **Observer notifications** via goroutines
3. **Error wrapping** with `wraperror.Errorf`

## Testing Conventions

- Tests use `test.Parallel()` for concurrent execution
- Use `require` for assertions that should fail the test
- Use `assert` for soft assertions
- Test helpers: `getTestObject(test)`, `printActual(test, actual)`
- Environment variable `SENZING_LOG_LEVEL` controls log verbosity

## Linting

Uses golangci-lint with 80+ rules. Key settings:

- JSON tags use UPPER_SNAKE_CASE
- Max line length: 120 characters (via golines)
- Coverage threshold: 75% (file, package, total)

## Related Repositories

- [sz-sdk-go](https://github.com/senzing-garage/sz-sdk-go) - Interface definitions
- [sz-sdk-go-core](https://github.com/senzing-garage/sz-sdk-go-core) - Native implementation
- [sz-sdk-go-grpc](https://github.com/senzing-garage/sz-sdk-go-grpc) - gRPC implementation

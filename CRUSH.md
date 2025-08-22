# goCBC Development Guide

Go port of the Python [caffeine-bedtime-calculator](https://github.com/sethll/caffeine-bedtime-calculator).
Calculates when substance levels (caffeine, nicotine) drop to target amounts for restful sleep using pharmacokinetic half-life modeling.

## Build/Test/Lint Commands
- **Build**: `make build` (includes git SHA) or `make build-dev` (without git SHA)
- **Build with git SHA**: `go build -ldflags "-X github.com/sethll/goCBC/pkg/progmeta.build=$(git rev-parse --short HEAD)" -o goCBC ./cmd/goCBC`
- **Run**: `make run` or `go run ./cmd/goCBC <target> '<time:amount>' ...`
- **Example**: `make run-example` or `go run ./cmd/goCBC 75 '1100:300' '1500:5'`
- **Test all**: `make test` or `go test ./...`
- **Test single package**: `make test-hlcalc` or `go test ./pkg/hlcalc`
- **Test with coverage**: `make test-coverage` or `go test -cover ./...`
- **Format**: `make fmt` or `go fmt ./...`
- **Vet**: `make vet` or `go vet ./...`
- **Mod tidy**: `make tidy` or `go mod tidy`
- **Check all**: `make check` (runs fmt, vet, test)
- **Clean**: `make clean`
- **Install**: `make install` (installs to GOPATH/bin)
- **Create tag**: `make tag VERSION=v1.0.0` (creates git tag for releases)
- **Help**: `make help`

## Usage
- **Target**: Milligrams of substance desired at bedtime (50-100 recommended for caffeine)
- **Time:Amount**: 24-hour format time and substance amount in mg (e.g., '1100:300')
- **Flags**:
  - `-c, --chem string`: Choose substance (default "caffeine")
  - `--list-chems`: List all available substance options with half-lives
  - `-v, --verbose`: Increase verbosity (use -v, -vv, -vvv, etc.)
- Supports multiple substance intake entries
- Outputs current substance level and anticipated bedtime
- **Available substances**: caffeine (5.7h half-life), nicotine (1.7h half-life)

## Code Style Guidelines
- **Package structure**: `cmd/` for binaries, `pkg/` for libraries
- **Imports**: Standard library first, then third-party, then local packages
- **Naming**: camelCase for variables/functions, PascalCase for exported items
- **Constants**: Use descriptive names like `CaffeineLambda`, `defaultTargetAmount`
- **Error handling**: Return errors explicitly, use slog for logging
- **Types**: Use descriptive struct names like `timeAndAmount`
- **Comments**: Package-level comments for exported functions, inline for complex logic
- **Logging**: Use slog with structured logging: `slog.Debug("message", "key", value)`

## Project Structure
- Main application logic in `cmd/goCBC/main.go`
- Calculation functions in `pkg/hlcalc/` (half-life calculations)
- Substance constants in `pkg/chems/` (caffeine: 5.7h, nicotine: 1.7h half-lives)
- Version/metadata in `pkg/progmeta/`
- Utility functions in `pkg/progutils/` (time ops, input validation, output formatting)
- Uses slog for structured logging
- Go 1.23.0 minimum version

## Dependencies
- **CLI Framework**: github.com/spf13/cobra (command-line interface)
- **Styling**: github.com/charmbracelet/lipgloss (terminal UI styling)
- **Logging**: github.com/charmbracelet/log (structured logging with slog integration)
- **Standard library**: Core Go packages for math, time, string operations
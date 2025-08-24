# goCBC Development Guide

Go port of the Python [caffeine-bedtime-calculator](https://github.com/sethll/caffeine-bedtime-calculator).
Calculates when substance levels (caffeine, nicotine) drop to target amounts for restful sleep using pharmacokinetic half-life modeling.

## Build/Test/Lint Commands
- **Build**: `make build` (includes git SHA)
- **Build with git SHA**: `go build -ldflags "-X github.com/sethll/goCBC/pkg/progmeta.build=$(git rev-parse --short HEAD)" -o ./build/goCBC ./cmd/goCBC`
- **Run**: `make run` or `go run ./cmd/goCBC <target> '<time:amount>' ...`
- **Example**: `make run-example` or `go run ./cmd/goCBC 75 '1100:300' '1500:150'`
- **Test all**: `make test` or `go test ./...` (not implemented yet)
- **Test single package**: `make test-hlcalc` or `go test ./pkg/hlcalc` (not implemented yet)
- **Test with coverage**: `make test-coverage` or `go test -cover ./...` (not implemented yet)
- **Format**: `make fmt` or `go fmt ./...`
- **Vet**: `make vet` or `go vet ./...`
- **Mod tidy**: `make tidy` or `go mod tidy`
- **Check all**: `make check` (runs fmt, vet, test)
- **Clean**: `make clean`
- **Install**: `make install` (installs to GOPATH/bin)
- **Create tag**: `make tag VERSION=v1.0.0` (creates git tag for releases)
- **Test version**: `make test-version` (shows current version tag)
- **Help**: `make help`

## Installation Methods
- **Go install**: `go install github.com/sethll/goCBC/cmd/goCBC@latest` (recommended)
- **Local build**: `make build` then use `./build/goCBC`
- **Development**: `go run ./cmd/goCBC <args>`

## Usage
- **Target**: Milligrams of substance desired at bedtime (50-100 recommended for caffeine)
- **Time:Amount**: 24-hour format time and substance amount in mg (e.g., '1100:300')
- **Flags**:
  - `-c, --chem string`: Choose substance (default "caffeine")
  - `--list-chems`: List all available substance options with half-lives
  - `--version`: Show version information
  - `-v, --verbose`: Increase verbosity (use -v, -vv, -vvv, etc.)
- Supports multiple substance intake entries
- Outputs current substance level and anticipated bedtime
- **Available substances**: caffeine (5.0h half-life), nicotine (4.25h half-life)

## Version Management
- **Single source of truth**: `pkg/progmeta/progmeta.go` contains definitive version (Major: "0", Minor: "1", Patch: "7")
- **Automated tag generation**: `cmd/version-tag` helper extracts version for git tagging
- **Version validation**: `init()` function validates numeric version components
- **Multi-scenario support**: Works with `go install`, local builds, and development builds
- **Fallback logic**: Build info → VCS info → ldflags → default

## Code Style Guidelines
- **Package structure**: `cmd/` for binaries, `pkg/` for libraries, `res/` for research documentation
- **Imports**: Standard library first, then third-party, then local packages
- **Naming**: camelCase for variables/functions, PascalCase for exported items
- **Constants**: Use descriptive names like `DefaultChem`, `CaffeineLambda`
- **Error handling**: Return errors explicitly, use slog for logging
- **Types**: Use descriptive struct names like `TimeAndAmount`
- **Documentation**: All exported items must have docstrings following Go conventions
- **Logging**: Use slog with structured logging: `slog.Debug("message", "key", value)`

## Project Structure
- **Main application**: `cmd/goCBC/main.go`
- **Version helper**: `cmd/version-tag/main.go` (extracts version for Makefile)
- **Calculation functions**: `pkg/hlcalc/` (half-life calculations with exponential decay)
- **Substance constants**: `pkg/chems/` (caffeine: 5.0h, nicotine: 4.25h half-lives)
- **Version/metadata**: `pkg/progmeta/` (centralized version management and program info)
- **Utility functions**: `pkg/progutils/` (time ops, input validation, output formatting)
- **Research documentation**: `res/` (pharmacokinetic analyses supporting substance parameters)
- Uses slog for structured logging with charmbracelet/log integration
- Go 1.23.0 minimum version

## Research-Based Parameters
- **Caffeine half-life**: 5.0 hours (updated from 4.7h based on comprehensive meta-analysis)
- **Nicotine half-life**: 4.25 hours (updated from 1.7h based on population pharmacokinetic studies)
- **Research documentation**: See `res/` directory for detailed pharmacokinetic analyses
- **Evidence-based**: All substance parameters backed by peer-reviewed literature

## Dependencies
- **CLI Framework**: github.com/spf13/cobra (command-line interface)
- **Styling**: github.com/charmbracelet/lipgloss (terminal UI styling and tables)
- **Logging**: github.com/charmbracelet/log (structured logging with slog integration)
- **Standard library**: Core Go packages for math, time, string operations, runtime/debug

## Development Workflow
1. **Make changes**: Edit code following style guidelines
2. **Format and check**: `make check` (runs fmt, vet, test)
3. **Test locally**: `make build && ./build/goCBC --version`
4. **Update version**: Edit `pkg/progmeta/progmeta.go` if needed
5. **Create tag**: `make tag VERSION=v0.1.8` (uses automated version extraction)
6. **Install test**: `go install github.com/sethll/goCBC/cmd/goCBC@latest`
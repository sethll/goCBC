# goCBC

A Go implementation of the [caffeine-bedtime-calculator](https://github.com/sethll/caffeine-bedtime-calculator) that calculates when substance levels (such as caffeine) drop to target amounts for restful sleep using pharmacokinetic half-life modeling.

## Description

goCBC helps you determine the optimal bedtime based on your substance intake throughout the day. By modeling the exponential decay of substances like caffeine using their known metabolic half-lives, the tool calculates when levels will drop to your desired target amount to improve sleep quality.

Supported substances:

```
# goCBC --list-chems
...
╭──────────────────────┬──────────────╮
│  Chem                │  Half-life   │
├──────────────────────┼──────────────┤
│  caffeine (default)  │  5.00 hours  │
│  nicotine            │  4.25 hours  │
╰──────────────────────┴──────────────╯
```

## Installation

### From Source

Requires Go 1.23.0 or later.

#### Easiest Method

```bash
go install github.com/sethll/goCBC/cmd/goCBC@latest
```

This method will not build the binary with the commit hash in version information, but it will be locked to the most recent tag. 

#### Other Easy Method

```bash
git clone https://github.com/sethll/goCBC.git
cd goCBC
make build
```

The binary will be created as `./build/goCBC` in the project directory. The binary will include the commit hash in version information. 

##### Install to GOPATH

```bash
make install
```

This installs the binary to `$GOPATH/bin/goCBC`.

## Usage

```
goCBC [flags] <target> '<time:amount>' ['<time:amount>' ...]
```

### Arguments

- **`target`**: Target substance level in milligrams at bedtime (50-100mg recommended for caffeine)
- **`'time:amount'`**: Substance intake in 24-hour format `'HHMM:amount'` (e.g., `'1100:300'` for 300mg at 11:00 AM)

### Flags

- `-c, --chem <substance>`: Choose substance (default: "caffeine")
- `-h, --help`: Show help information
- `--list-chems`: List all available substances with their half-lives
- `-v, --verbose`: Increase verbosity (use `-v`, `-vv`, `-vvv` for more detail)
- `--version`: Show version information 

### Examples

Calculate bedtime with caffeine intake:
```bash
goCBC 50 '1100:300' '1500:150'
# equivalent to 
# goCBC --chem caffeine 50 '1100:300' '1500:150'
```

Calculate with nicotine:
```bash
goCBC --chem nicotine 0.2 '0900:2' '1300:4'
```

List available substances:
```bash
goCBC --list-chems
```

### Sample Output

```
 Caffeine remaining in system:     ~86mg
 Reach target (50mg) for sleep at: 2025-08-23 03:30
```

## Development

### Building

```bash
make build          # Build with git SHA
make build-dev      # Build without git SHA
```

### Code Quality

```bash
make check          # Run fmt, vet 
make fmt            # Format code
make vet            # Run go vet
```

## License

This project is licensed under the GPLv3 License - see the [LICENSE](LICENSE) file for details.

## Related Projects

- [caffeine-bedtime-calculator](https://github.com/sethll/caffeine-bedtime-calculator) - Original Python implementation
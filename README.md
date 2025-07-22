# My Tools

A collection of command-line utilities written in Go using Cobra CLI framework.

## Tools

### filter
A powerful text filtering tool that supports both file input and stdin piping.

**Features:**
- Filter lines containing specific patterns
- Read from files or stdin
- Verbose output with match indicators
- Built-in help system

**Usage Examples:**

```bash
# Basic filtering from stdin
echo -e "info: starting\nerror: failed\ninfo: done" | filter --pattern "error"

# Filter from a file
filter --file logs.txt --pattern "WARNING"

# Verbose output with match indicators
cat application.log | filter --pattern "ERROR" --verbose

# Show all lines (no pattern filter)
filter --file input.txt

# Get help
filter --help
filter -h
```

**Command-line Options:**
- `-f, --file`: Input file (default: read from stdin)
- `-p, --pattern`: Pattern to filter by (shows all lines if empty)
- `-v, --verbose`: Verbose output with match indicators
- `-h, --help`: Show help message

## Building

### Prerequisites
- Go 1.24 or later

### Build locally
```bash
# Download dependencies
go mod download

# Build the filter tool
mkdir -p bin
go build -o bin/filter ./cmd/filter

# Test it works
echo "test line" | ./bin/filter --pattern "test"
```

### Install from source
```bash
# Install directly from GitHub
go install github.com/aryan-binazir/tools/cmd/filter@latest

# Or clone and install locally
git clone https://github.com/aryan-binazir/tools.git
cd tools
go install ./cmd/filter
```

## Development

### Project Structure
```
├── cmd/                    # Command-line applications
│   └── filter/            # Filter tool
│       ├── main.go        # Entry point
│       └── cmd/           # Cobra commands
│           └── root.go    # Root command definition
├── internal/              # Private application code
│   └── shared/            # Shared utilities
│       └── utils.go       # File/stdin handling
├── .github/workflows/     # GitHub Actions CI/CD
└── README.md
```

### Running Tests
```bash
go test -v ./...
```

### Code Quality
```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Run linter (if golangci-lint is installed)
golangci-lint run
```

## Examples

### Real-world Usage Scenarios

```bash
# Monitor log files for errors
tail -f /var/log/app.log | filter --pattern "ERROR" --verbose

# Extract warnings from multiple log files
cat *.log | filter --pattern "WARN"

# Filter configuration files for specific settings
filter --file config.ini --pattern "database"

# Process CSV data (filter rows containing specific values)
filter --file data.csv --pattern "production"

# Chain with other tools
grep -r "TODO" . | filter --pattern "urgent" --verbose
```

### Integration with Other Tools

```bash
# With find and xargs
find . -name "*.log" -exec cat {} \; | filter --pattern "critical"

# With curl and jq (for JSON logs)
curl -s https://api.example.com/logs | jq -r '.[]' | filter --pattern "error"

# With systemd journal
journalctl -f | filter --pattern "authentication"
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.
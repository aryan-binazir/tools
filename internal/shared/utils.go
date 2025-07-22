package shared

import (
	"bufio"
	"io"
	"os"
)

// GetInputReader determines where to read input from
// Returns the reader, a cleanup function, and any error
func GetInputReader(filename string) (io.Reader, func() error, error) {
	if filename != "" {
		// Read from file
		file, err := os.Open(filename)
		if err != nil {
			return nil, nil, err
		}
		return file, file.Close, nil
	}

	// Read from stdin
	return os.Stdin, func() error { return nil }, nil
}

// ProcessLines reads lines from reader and calls processor function
func ProcessLines(reader io.Reader, processor func(string) string) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		processed := processor(line)
		if processed != "" {
			println(processed)
		}
	}
	return scanner.Err()
}

// GetVersion returns the current version of the tools
func GetVersion() string {
	return "v0.1.0"
}

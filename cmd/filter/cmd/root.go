package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/aryan-binazir/tools/internal/shared"
	"github.com/spf13/cobra"
)

var (
	inputFile string
	pattern   string
	verbose   bool
)

var rootCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filter text based on patterns",
	Long: `Filter tool reads from stdin or a file and filters lines based on patterns.

The filter tool supports both command-line arguments and piped input, making it
flexible for various text processing workflows.`,

	Example: `  # Filter lines from stdin containing "error"
  echo -e "info: starting\nerror: failed\ninfo: done" | filter --pattern "error"
  
  # Filter lines from a file
  filter --file input.txt --pattern "warning"
  
  # Verbose output with match indicators
  cat log.txt | filter --pattern "ERROR" --verbose
  
  # Show help
  filter --help`,

	RunE: func(cmd *cobra.Command, args []string) error {
		reader, closer, err := shared.GetInputReader(inputFile)
		if err != nil {
			return fmt.Errorf("failed to open input: %w", err)
		}
		defer closer()

		if verbose {
			fmt.Fprintf(os.Stderr, "Hello from filter v%s - filtering with pattern: %s\n",
				shared.GetVersion(), pattern)
		}

		return shared.ProcessLines(reader, func(line string) string {
			if pattern == "" || strings.Contains(line, pattern) {
				if verbose {
					return fmt.Sprintf("[MATCH] %s", line)
				}
				return line
			}
			return ""
		})
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "file", "f", "", "Input file (default: read from stdin)")
	rootCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Pattern to filter by (shows all lines if empty)")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output with match indicators")

	// Cobra automatically handles --help and -h flags
}

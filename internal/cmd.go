// Provides core logic for podlet2nix for parsing podman quadlet unit files
// and converting them to Nix configuration format.
package internal

import (
	"fmt"
	"io"
	"os"
)

// GetReader returns an io.ReadCloser for reading input data.
// If no command-line arguments are provided, it reads from stdin.
// If one argument is provided, it treats it as a file path.
func GetReader() io.ReadCloser {
	var reader io.ReadCloser
	var err error

	if len(os.Args) == 1 {
		reader = io.NopCloser(os.Stdin)
	} else {
		path := os.Args[1]

		reader, err = os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Could not open file '%s': %v\n", path, err)
			os.Exit(1)
		}
	}

	return reader
}

package internal

import (
	"fmt"
	"io"
	"os"
)

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

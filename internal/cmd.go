package internal

import (
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
			panic(err)
		}
	}

	return reader
}

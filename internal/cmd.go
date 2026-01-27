package internal

import (
	"io"
	"os"
)

func GetReader() io.Reader {
	var reader io.Reader
	var err error

	if len(os.Args) == 1 {
		reader = os.Stdin
	} else {
		path := os.Args[1]

		reader, err = os.Open(path)
		if err != nil {
			panic(err)
		}
	}

	return reader
}

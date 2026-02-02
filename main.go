package main

import (
	"github.com/thilobillerbeck/podlet2nix/internal"
)

func main() {
	reader := internal.GetReader()
	defer reader.Close()
	internal.ParseReader(reader)
}

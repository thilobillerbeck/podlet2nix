package main

import (
	"github.com/thilobillerbeck/podlet2nix/internal"
)

func main() {
	internal.ParseReader(internal.GetReader())
}

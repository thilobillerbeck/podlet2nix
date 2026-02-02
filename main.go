package main

import (
	"fmt"
	"os"

	"github.com/thilobillerbeck/podlet2nix/internal"
)

var (
	Version = "dev"
)

func preParseFlags(flag string) {
	if flag == "--version" {
		fmt.Println(Version)
		os.Exit(0)
	}
	if flag == "--help" {
		fmt.Println("Usage with files: podlet2nix FILE")
		fmt.Println("Usage with stdin: podlet [...] | podlet2nix")
		fmt.Println("")
		fmt.Println("  --version: Print version")
		fmt.Println("  --help: Print this help message")
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) > 1 {
		preParseFlags(os.Args[1])
	}

	reader := internal.GetReader()
	defer reader.Close()
	internal.ParseReader(reader)
}

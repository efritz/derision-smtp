package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
)

func main() {
	kingpin.Parse()

	if err := serve(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}

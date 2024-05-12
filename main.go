package main

import (
	"fmt"
	"os"
)

func main() {
	flagParse()

	switch {
	case help:
		usage()
		os.Exit(0)
	}

	if err := run(); err != nil {
		fmt.Printf("errors: %s", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Hello world")
	return nil
}

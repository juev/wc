package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	bytes = flag.Bool(
		"c",
		false,
		`The number of bytes in each input file is written to the standard output.
		This will cancel out any prior usage of the -m option.`)

	lines = flag.Bool(
		"l",
		false,
		`The number of lines in each input file is written to the standard output.`,
	)
	chars = flag.Bool(
		"m",
		false,
		`The number of characters in each input file is written to the standard output.
		If the current locale does not support multibyte characters, this is equivalent to
		the -c option.  This will cancel out any prior usage of the -c option.`)
	words = flag.Bool(
		"w",
		false,
		`The number of words in each input file is written to the standard output.`)
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Printf("errors: %s", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Hello world")
	return nil
}

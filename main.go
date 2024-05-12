package main

import (
	"bufio"
	"fmt"
	"io"
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
	if len(files) == 0 {
		c, s, l, _ := processFile()
		fmt.Println(c, s, l)
	}

	for i := range files {
		c, s, l, _ := processFile(files[i])
		fmt.Println(c, s, l, files[i])
	}

	return nil
}

func processFile(fileName ...string) (count, size, lines int64, err error) {
	file := os.Stdin
	if len(fileName) != 0 {
		file, err = os.Open(fileName[0])
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		r, s, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, 0, err
		}
		if r == '\n' {
			lines++
		}
		count++
		size += int64(s)
	}

	return count, size, lines, nil
}

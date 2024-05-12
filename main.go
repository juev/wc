package main

import (
	"bufio"
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
	if len(files) == 0 {
		c, _ := processFile()
		fmt.Println("bytes:", c)
	}

	for i := range files {
		c, _ := processFile(files[i])
		fmt.Println(c, files[i])
	}

	return nil
}

func processFile(fileName ...string) (count int64, err error) {
	file := os.Stdin
	if len(fileName) != 0 {
		file, err = os.Open(fileName[0])
		if err != nil {
			panic(err)
		}
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		count++
	}

	file.Close()

	return count, nil
}

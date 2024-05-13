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
		c, s, l, err := processFile()
		if err != nil {
			return err
		}
		fmt.Println(c, s, l)
		return nil
	}

	type ff struct {
		Count int64
		Size  int64
		Lines int64
		File  string
	}
	var FF = make([]ff, 0, len(files))
	var totalCount, totalSize, totalLines int64
	for i := range files {
		c, s, l, _ := processFile(files[i])
		FF = append(FF, ff{
			Count: c,
			Size:  s,
			Lines: l,
			File:  files[i],
		})
		totalCount += c
		totalSize += s
		totalLines += l
	}

	for i := range FF {
		fmt.Println(FF[i].Count, FF[i].Size, FF[i].Lines, FF[i].File)
	}
	if len(FF) > 1 {
		fmt.Println(totalCount, totalSize, totalLines)
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

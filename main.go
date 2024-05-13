package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
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
		c, s, l, w, err := processFile()
		if err != nil {
			return err
		}
		print(c, s, l, w, "")
		return nil
	}

	type ff struct {
		Count int64
		Size  int64
		Lines int64
		Words int64
		File  string
	}
	var FF = make([]ff, 0, len(files))
	var totalCount, totalSize, totalLines, totalWords int64
	for i := range files {
		c, s, l, w, err := processFile(files[i])
		if err != nil {
			return err
		}
		FF = append(FF, ff{
			Count: c,
			Size:  s,
			Lines: l,
			Words: w,
			File:  files[i],
		})
		totalCount += c
		totalSize += s
		totalLines += l
		totalWords += w
	}

	for i := range FF {
		print(FF[i].Count, FF[i].Size, FF[i].Lines, FF[i].Words, FF[i].File)
	}

	if len(FF) > 1 {
		print(totalCount, totalSize, totalLines, totalWords, "total")
	}

	return nil
}

func processFile(fileName ...string) (count, size, lines, words int64, err error) {
	file := os.Stdin
	if len(fileName) != 0 {
		file, err = os.Open(fileName[0])
		if err != nil {
			return 0, 0, 0, 0, nil
		}
	}
	defer file.Close()

	var prev rune
	reader := bufio.NewReader(file)
	for {
		r, s, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, 0, 0, err
		}
		if r == '\n' {
			lines++
		}
		if !unicode.IsSpace(r) && unicode.IsSpace(prev) {
			words++
		}
		prev = r
		count++
		size += int64(s)
	}

	return count, size, lines, words, nil
}

func print(c, s, l, w int64, file string) {
	if chars {
		fmt.Print(c, "\t")

	}
	if bytes {
		fmt.Print(s, "\t")
	}
	if lines {
		fmt.Print(l, "\t")
	}
	if words {
		fmt.Print(w, "\t")
	}
	fmt.Println(file)
}

package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

var (
	bytes, lines, chars, words, help bool
	files                            []string
)

func flagParse() {
	flag.BoolVarP(
		&bytes,
		"bytes",
		"c",
		false,
		`The number of bytes in each input file is written to the
		standard output.`)

	flag.BoolVarP(
		&lines,
		"lines",
		"l",
		false,
		`The number of lines in each input file is written to the
		standard output.`,
	)
	flag.BoolVarP(
		&chars,
		"chars",
		"m",
		false,
		`The number of characters in each input file is written to the standard output.`)
	flag.BoolVarP(
		&words,
		"words",
		"w",
		false,
		`The number of words in each input file is written to the standard output.`)

	flag.BoolVarP(&help, "help", "h", false, "help string")

	flag.Parse()
	files = flag.Args()
}

func usage() {
	fmt.Print(`
NAME
     wc – word, line, character, and byte count

SYNOPSIS
     wc [--libxo] [-Lclmw] [file ...]

DESCRIPTION
    The wc utility displays the number of lines, words, and bytes contained in
    each input file, or standard input (if no file is specified) to the
    standard output.  A line is defined as a string of characters delimited by
    a ⟨newline⟩ character.  Characters beyond the final ⟨newline⟩ character will
    not be included in the line count.

    A word is defined as a string of characters delimited by white space
    characters.  White space characters are the set of characters for which the
    iswspace(3) function returns true.  If more than one input file is
    specified, a line of cumulative counts for all the files is displayed on a
    separate line after the output for the last file.

    The following options are available:

    -c      The number of bytes in each input file is written to the standard
    output.  This will cancel out any prior usage of the -m option.

    -l      The number of lines in each input file is written to the standard
    output.

    -m      The number of characters in each input file is written to the
    standard output.  If the current locale does not support multibyte
    characters, this is equivalent to the -c option.  This will cancel out any
    prior usage of the -c option.

    -w      The number of words in each input file is written to the standard
    output.

    When an option is specified, wc only reports the information requested by
    that option.  The order of output always takes the form of line, word,
    byte, and file name.  The default action is equivalent to specifying the
    -c, -l and -w options.

    If no files are specified, the standard input is used and no file name is
    displayed.  The prompt will accept input until receiving EOF, or [^D] in
    most environments.

    If wc receives a SIGINFO (see the status argument for stty(1)) signal, the
    interim data will be written to the standard error output in the same
    format as the standard completion message.

ENVIRONMENT
    The LANG, LC_ALL and LC_CTYPE environment variables affect the execution of
    wc as described in environ(7).

EXIT STATUS
    The wc utility exits 0 on success, and >0 if an error occurs.

EXAMPLES
    Count the number of characters, words and lines in each of the files
    report1 and report2 as well as the totals for both:

        wc -mlw report1 report2

    Find the longest line in a list of files:

        wc -L file1 file2 file3 | fgrep total
`)
}

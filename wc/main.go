package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"github.com/mattn/go-isatty"
)

func main() {
	flag.Parse()

	if !isatty.IsTerminal(os.Stdin.Fd()) {
		nlines, nwords, nchars := count(os.Stdin)
		fmt.Printf("%d %d %d\n", nlines, nwords, nchars)
		os.Exit(0)
	}

	var totalLines, totalWords, totalChars int
	for _, filepath := range flag.Args() {
		func(filepath string) {
			file, err := os.Open(filepath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "wc: %s: open: No such file or directory\n", filepath)
				return
			}
			defer file.Close()
			nlines, nwords, nchars := count(file)
			fmt.Printf("%d %d %d %s\n", nlines, nwords, nchars, filepath)
			totalLines += nlines
			totalWords += nwords
			totalChars += nchars
		}(filepath)
	}

	if len(flag.Args()) > 1 {
		fmt.Printf("%d %d %d total\n", totalLines, totalWords, totalChars)
	}
}

func count(in io.Reader) (nlines, nwords, nchars int) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		nlines++
		line := scanner.Text()
		nwords += len(strings.FieldsFunc(line, func(c rune) bool {
			return unicode.IsSpace(c)
		}))
		nchars += len(line) + 1 // +1 for \n
	}
	return nlines, nwords, nchars
}

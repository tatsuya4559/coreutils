package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

var (
	l = flag.Bool("l", false, "count lines")
	w = flag.Bool("w", false, "count words")
	c = flag.Bool("c", false, "count chars")
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		nlines, nwords, nchars := count(os.Stdin)
		printCount(nlines, nwords, nchars, "")
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
			printCount(nlines, nwords, nchars, filepath)
			totalLines += nlines
			totalWords += nwords
			totalChars += nchars
		}(filepath)
	}

	if len(flag.Args()) > 1 {
		printCount(totalLines, totalWords, totalChars, "total")
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

func printCount(nlines, nwords, nchars int, filepath string) {
	// どのフラグも立たない場合は全部
	if !(*l || *w || *c) {
		*l = true
		*w = true
		*c = true
	}
	var b strings.Builder
	if *l {
		fmt.Fprintf(&b, "	%d", nlines)
	}
	if *w {
		fmt.Fprintf(&b, "	%d", nwords)
	}
	if *c {
		fmt.Fprintf(&b, "	%d", nchars)
	}
	if filepath != "" {
		fmt.Fprintf(&b, "	%s", filepath)
	}
	fmt.Println(b.String())
}

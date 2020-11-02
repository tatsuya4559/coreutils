package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	PAGELEN = 24
	LINELEN = 512
)

func main() {
	if len(os.Args) == 1 {
		doMore(os.Stdin)
	} else {
		for _, arg := range os.Args[1:] {
			fp, err := os.Open(arg)
			if err != nil {
				panic(err)
			}
			doMore(fp)
			fp.Close()
		}
	}
}

func doMore(r io.Reader) {
	tty, err := os.Open("/dev/tty")
	if err != nil {
		panic(err)
	}

	var numLines int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if numLines == PAGELEN {
			reply := seeMore(tty)
			if reply == 0 {
				break
			}
			numLines -= reply
		}

		fmt.Println(scanner.Text())
		numLines++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func seeMore(r io.Reader) int {
	fmt.Printf("\033[7m more? \033[m")

	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		panic(err)
	}

	switch string(buf) {
	case "q":
		return 0
	case " ":
		return PAGELEN
	case "\n":
		return 1
	default:
		return 0
	}
}

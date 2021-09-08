package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const BUFSIZE = 4096

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		cat(os.Stdin)
	} else {
		for _, filepath := range flag.Args() {
			func(filepath string) {
				file, err := os.Open(filepath)
				if err != nil {
					fmt.Fprintf(os.Stderr, "cat: %s: No such file or directory", filepath)
				}
				defer file.Close()
				cat(file)
			}(filepath)
		}
	}
}

func cat(in io.Reader) {
	buf := make([]byte, BUFSIZE)
	for {
		n, err := in.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			log.Fatal(err)
		}
		if n == 0 {
			return
		}
		n, err = os.Stdout.Write(buf[:n])
		if err != nil && !errors.Is(err, io.EOF) {
			log.Fatal(err)
		}
		if n == 0 {
			log.Fatal("write 0 byte")
		}
	}
}

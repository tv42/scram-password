package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"eagain.net/go/scram-password/internal/scramble"
)

func run(usernames []string) error {
	scanner := bufio.NewScanner(os.Stdin)
	for _, username := range usernames {
		if !scanner.Scan() {
			// error or EOF
			err := scanner.Err()
			if err == nil {
				err = io.ErrUnexpectedEOF
			}
			return fmt.Errorf("reading passwords from standard input: %w", err)
		}
		password := scanner.Text()
		salted, err := scramble.Hash(username, password)
		if err != nil {
			return fmt.Errorf("cannot SCRAM hash: %w", err)
		}
		fmt.Println(salted)
	}
	return nil
}

var prog = filepath.Base(os.Args[0])

func usage() {
	out := flag.CommandLine.Output()
	fmt.Fprintf(out, "Usage: %s [OPTS] USERNAME.. <PASSWORDS\n", prog)
	fmt.Fprintf(out, "Outputs a SCRAM hashed password.\n")
	fmt.Fprintf(out, "\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		log.Printf("missing username")
		os.Exit(2)
	}

	if err := run(flag.Args()); err != nil {
		log.Fatalf("error:: %v", err)
	}
}

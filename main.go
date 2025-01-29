package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ExitCodeOK = iota
	ExitCodeError
)

var commandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

// run is the entry point for the csv-split command.
func run(args []string) int {
	recordCount := commandLine.Int("record_count", 100, "Create split files record_count records.")
	if err := commandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse flags: %v\n", err)
		return ExitCodeError
	}

	if commandLine.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Only one file can be processed at a time. Usage: csv-split <file>")
		return ExitCodeError
	}
	fileName := commandLine.Arg(0)

	fmt.Printf("record_count: %v\n", *recordCount)
	fmt.Printf("file_name: %v\n", fileName)

	return ExitCodeOK
}

func main() {
	os.Exit(run(os.Args[1:]))
}

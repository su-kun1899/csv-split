package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: csv-split <input-file> <lines-per-file>")
		return
	}
	fmt.Println("CSV Splitter CLI")
}

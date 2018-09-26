package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	const (
		INDEX  = "index"
		SEARCH = "search"
	)
	// Create the gosearch command line object.
	cmd := NewSearchCMD()
	// Check to see if enough arguments were given to the command.
	if len(os.Args) < 2 {
		fmt.Println("Not enough commands supplied.")
		return
	}
	path := flag.String("path", "default", "Index a file to search later.")
	flag.Parse()
	// Now check to see what type of command has been supplied to gosearch.
	if os.Args[1] == INDEX {
		cmd.IndexCMD(*path)
	}
}

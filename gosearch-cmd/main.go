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
	path := flag.String("path", "test.txt", "Index a file to search later.")
	key := flag.String("key", "default", "The word that you are searching for.")
	flag.Parse()
	// Now check to see what type of command has been supplied to gosearch.
	cmd.IndexCMD(*path)
	if *key != "" {
		cmd.SearchCMD(*key)
	}
}

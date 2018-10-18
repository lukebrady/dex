package main

import (
	"flag"
	"fmt"
)

func main() {
	const (
		INDEX  = "index"
		SEARCH = "search"
	)
	cmd := NewSearchCMD()
	// Index a file and serialize its inverted index to disk.
	index := flag.String("index", "default", "Supply the path to the file you would like to index.")
	search := flag.String("search", "default", "The word that you are searching for out of all indexed files.")
	file := flag.String("file", "default", "Index all of the files within an index.conf file.")
	flag.Parse()
	// Now check to see what type of command has been supplied.
	if *index != "default" {
		cmd.IndexCMD(*index)
	} else if *search != "default" {
		cmd.SearchCMD(*search)
	} else if *file != "default" {
		cmd.FileCMD(*file)
	} else {
		fmt.Printf("Incorrect use of the gosearch tool.\n")
	}
}

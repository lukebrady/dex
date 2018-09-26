package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/fatih/color"
)

// InvertedIndex struct
type InvertedIndex struct {
	index map[string]*ValueNode
	mutex *sync.Mutex
	size  uint
}

// ValueNode struct
type ValueNode struct {
	value string
	index int
	next  *ValueNode
}

// NewIndex returns a pointer to an Inverted Index object.
func NewIndex() *InvertedIndex {
	// Make a new map that can be given to the InvertedIndex.
	ind := make(map[string]*ValueNode)
	return &InvertedIndex{
		index: ind,
		mutex: &sync.Mutex{},
		size:  0,
	}
}

// IndexFile reads a file and indexes it.
func (i *InvertedIndex) IndexFile(file string) error {
	// Read the given file into memory. This should be changed in the future.
	fopen, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	// Create a Regex object that can match punctuation within a file's text.
	re, err := regexp.Compile("[.,?]")
	if err != nil {
		return err
	}
	// Remove all punctuation from the file.
	premove := re.ReplaceAllString(string(fopen), " ")
	// Split all of the strings within the file.
	str := strings.Split(premove, " ")
	// Now sort all of the words within the file.
	sort.Strings(str)
	// Enter all of the values found in the file into the index.
	i.mutex.Lock()
	for index, word := range str {
		if i.index[word] == nil {
			val := &ValueNode{
				value: file,
				index: index,
				next:  nil,
			}
			i.index[word] = val
			// Increase the size count of the index.
			i.size++
		} else {
			// Create the value node that will be inserted into the chain.
			val := &ValueNode{
				value: file,
				next:  nil,
			}
			// Assign the root value to the first value node.
			place := i.index[word]
			for place.next != nil {
				place = place.next
			}
			// Place the value node when .next == nil.
			place.next = val
			// Increase size of the total inverse index.
			i.size++
		}
	}
	i.mutex.Unlock()
	// Return nil if no error occurs.
	return nil
}

// SearchByKey searches all indexed documents for the provided key and prints where
// the word occurs within that document.
func (i *InvertedIndex) SearchByKey(key string) {
	if i.index[key] != nil {
		// Assign the root value to the first value node.
		place := i.index[key]
		for place != nil {
			document, err := ioutil.ReadFile(place.value)
			if err != nil {
				println(err)
			}
			str := strings.Split(string(document), " ")
			// Create a color print function.
			cyan := color.New(color.FgCyan).PrintfFunc()
			// Print out the document where the word was found.
			fmt.Printf("\"%s\" found in %s.\n", key, place.value)
			// After Reading in the document, print to STDOUT.
			for _, keyword := range str {
				// If the keyword matches the key, print the word out.
				if key == keyword {
					cyan(" %s ", keyword)
				} else {
					fmt.Printf(" %s ", keyword)
				}
			}
			fmt.Println()
			place = place.next
		}
		// Increase size of the total inverse index.
		i.size++
	} else {
		fmt.Println("This key has no entries in the index.")
	}
}

// PrintIndex prints the given key's index.
func (i *InvertedIndex) PrintIndex() {
	// Print the entire index. Will only print keys and ValueNode address.
	fmt.Println(i.index)
}

// PrintByKey prints a key's entire chain.
func (i *InvertedIndex) PrintByKey(key string) {
	fmt.Printf("%s:\n", key)
	if i.index[key] != nil {
		place := i.index[key]
		for place != nil {
			fmt.Printf("%s\n", place.value)
			place = place.next
		}
	} else {
		fmt.Println("This key has no entries in the index.")
	}
}

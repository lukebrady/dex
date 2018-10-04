package main

import "encoding/gob"
import "bytes"
import "os"

import "path/filepath"

// GoSearchCMD is the command line object that is used to control the gosearch utility.
type GoSearchCMD struct {
	Index *InvertedIndex
}

// NewSearchCMD returns a new GoSearchCMD object that will be used to track command line inputs.
func NewSearchCMD() *GoSearchCMD {
	ii := NewIndex()
	return &GoSearchCMD{
		Index: ii,
	}
}

// EncodeMap takes a token map and serializes it to disk for use by nigel-bot.
// This will be used to generate testing Maps that will eventually be replaced by a database.
func EncodeMap(tokenMap *InvertedIndex, encodedIndex chan []byte) {
	buf := &bytes.Buffer{}
	// Create an encoder to serialize the map.
	encoder := gob.NewEncoder(buf)
	// Now encode the map in gob format.
	if err := encoder.Encode(tokenMap); err != nil {
		panic(err)
	}
	// Now return the encoded map to be written to disk.
	encodedIndex <- buf.Bytes()
}

// IndexCMD is the function that indexes a file into memory that can be searched.
func (cmd *GoSearchCMD) IndexCMD(path string) {
	// Get the filename from the filepath.
	fileName := filepath.Base(path)
	// Now run the cmd.Index.IndexFile(path) function to index the file path supplied to the command.
	cmd.Index.IndexFile(path)
	// Make the channel to retrieve encoded data.
	byteChan := make(chan []byte)
	// Now serialize the new index to disk so that it can be used later to search.
	go EncodeMap(cmd.Index, byteChan)
	// Now write to disk.
	writePath := "dex/tmp/" + fileName + ".gob"
	indexFile, err := os.OpenFile(writePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer indexFile.Close()
	_, err = indexFile.Write(<-byteChan)
}

// SearchCMD will search for the supplied token or tokens in a group of words.
func (cmd *GoSearchCMD) SearchCMD(key string) {
	err := cmd.Index.SearchByKey(key)
	if err != nil {
		println(err)
	}
}

// FileCMD will index all of the files within a supplied configuration file.
func (cmd *GoSearchCMD) FileCMD(path string) {
	// Get the index supplied index configuration.
	indexConf := NewIndexConfigurationObject(path)
	for _, file := range indexConf.IndexFiles {
		// Get the filename from the filepath.
		fileName := filepath.Base(file)
		// Now run the cmd.Index.IndexFile(path) function to index the file path supplied to the command.
		cmd.Index.IndexFile(file)
		// Make the channel to retrieve encoded data.
		byteChan := make(chan []byte)
		// Now serialize the new index to disk so that it can be used later to search.
		go EncodeMap(cmd.Index, byteChan)
		// Now write to disk.
		writePath := "dex/tmp/" + fileName + ".gob"
		indexFile, err := os.OpenFile(writePath, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer indexFile.Close()
		_, err = indexFile.Write(<-byteChan)
	}
}

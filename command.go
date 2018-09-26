package main

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

// IndexCMD is the function that indexes a file into memory that can be searched.
func (cmd *GoSearchCMD) IndexCMD(path string) {
	// Now run the cmd.Index.IndexFile(path) function to index the file path supplied to the command.
	err := cmd.Index.IndexFile(path)
	if err != nil {
		panic(err)
	}
}

// SearchCMD will search for the supplied token or tokens in a group of words.
func (cmd *GoSearchCMD) SearchCMD(key string) {
	err := cmd.Index.SearchByKey(key)
	if err != nil {
		println(err)
	}
}

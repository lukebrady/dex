package main

// InvertedIndex struct
type InvertedIndex struct {
	Index map[string]*ValueNode
	Size  uint
}

// ValueNode struct
type ValueNode struct {
	Value string
	Index int
	Next  *ValueNode
}

// GoSearchCMD is the command line object that is used to control the gosearch utility.
type GoSearchCMD struct {
	Index *InvertedIndex
}

// SearchConfiguration is an object that will be used to tune the gosearch program.
type SearchConfiguration struct {
	OutputColor string `json:"OutputColor"`
}

// IndexFile is a type that contains a list of the files to index.
type IndexFile struct {
	IndexFiles []string `json:"Index_Files"`
}

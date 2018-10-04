package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fatih/color"
)

// SearchConfiguration is an object that will be used to tune the gosearch program.
type SearchConfiguration struct {
	OutputColor string `json:"OutputColor"`
}

// NewConfigurationObject returns a config object that will be used to tune gosearch.
func NewConfigurationObject() *SearchConfiguration {
	// Read the in the binary data from the config file.
	jsonConf, err := ioutil.ReadFile("dex/config/gosearch.conf")
	if err != nil {
		panic(err)
	}
	conf := &SearchConfiguration{}
	// Now unmarshal the JSON to the SearchConfiguration object.
	if err := json.Unmarshal(jsonConf, conf); err != nil {
		panic(err)
	}
	// Return the configuration object.
	return conf
}

// NewIndexConfigurationObject returns a config object that will be used to tune gosearch.
func NewIndexConfigurationObject(path string) *IndexFile {
	// Read the in the binary data from the config file.
	jsonConf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	conf := &IndexFile{}
	// Now unmarshal the JSON to the SearchConfiguration object.
	if err := json.Unmarshal(jsonConf, conf); err != nil {
		panic(err)
	}
	// Return the configuration object.
	return conf
}

// GetColor takes the color data from the gosearch.conf file and translates it into a color
// that can be used by the program to determine FGColor.
func (s *SearchConfiguration) GetColor() color.Attribute {
	if s.OutputColor == "Cyan" {
		return color.FgCyan
	} else if s.OutputColor == "Blue" {
		return color.FgBlue
	} else if s.OutputColor == "Green" {
		return color.FgHiGreen
	}
	return color.FgHiRed
}

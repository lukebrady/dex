package main

import (
	"encoding/json"
	"io/ioutil"
)

// SearchConfiguration is an object that will be used to tune the gosearch program.
type SearchConfiguration struct {
	OutputColor string `json:"OutputColor"`
}

// NewConfigurationObject returns a config object that will be used to tune gosearch.
func NewConfigurationObject() *SearchConfiguration {
	// Read the in the binary data from the config file.
	jsonConf, err := ioutil.ReadFile("../config/gosearch.conf")
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

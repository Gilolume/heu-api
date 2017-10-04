package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var configFile = "config.json"
var config Configuration

func readConfigFile(filename string) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/flw-cn/go-study/oo/base"
	"github.com/flw-cn/go-study/oo/greet"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Debug   bool         `json:"Debug" yaml:"Debug" flag:"|false|Debug"`
	LogFile string       `json:"LogFile" yaml:"LogFile" flag:"||LogFile"`
	Greet   greet.Config `json:"Greet" yaml:"Greet" flag:"||Greet"`
}

func main() {
	var config Config
	var err error

	if len(os.Args) > 1 && os.Args[1] == "yaml" {
		err = yamlConfig("config.yaml", &config)
	} else if len(os.Args) > 1 && os.Args[1] == "json" {
		err = jsonConfig("config.json", &config)
	} else {
		log.Printf("Usage: progName <yaml|json>")
		os.Exit(1)
	}
	log.Printf("config: %#v", config)

	if config.Greet.LogFile == "" {
		config.Greet.LogFile = config.LogFile
	}

	var f base.Service
	f = greet.NewGreet(config.Greet)

	f.Init()

	err = f.Start()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	debug := false
	for {
		select {
		case <-time.After(3 * time.Second):
			debug = !debug
			f.SetDebug(debug)
		}
	}
}

func jsonConfig(file string, config interface{}) error {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(text, config)
}

func yamlConfig(file string, config interface{}) error {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(text, config)
}

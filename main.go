package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/xJustJqy/MockingJay/server"
	"github.com/xJustJqy/MockingJay/server/player/chat"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	config, err := readConfig()
	if err != nil {
		log.Fatalln(err)
	}

	srv := server.New(&config, log)
	srv.CloseOnProgramEnd()
	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}

	for srv.Accept(nil) {
	}
}

// readConfig reads the configuration from the config.json file, or creates the file if it does not yet exist.
func readConfig() (server.Config, error) {
	c := server.DefaultConfig()
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		data, err := json.Marshal(c.Marshalable())
		if err != nil {
			return c, fmt.Errorf("failed encoding default config: %v", err)
		}
		if err := ioutil.WriteFile("config.json", data, 0644); err != nil {
			return c, fmt.Errorf("failed creating config: %v", err)
		}
		return c, nil
	}
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return c, fmt.Errorf("error reading config: %v", err)
	}
	if err := json.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("error decoding config: %v", err)
	}
	return c, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Logging bool    `json:"consolelogging"`
	Game    Game    `json:"game"`
	Discord Discord `json:"discord"`
	Telnet  Telnet  `json:"telnet"`
}

type Game struct {
	BloodMoonFrequency int `json:"bloodmoonfrequency"`
}

type Discord struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
	Prefix  string `json:"prefix"`
}

type Telnet struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func ParseConfiguration() Configuration {
	// Parse config json
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var c Configuration

	// parse json
	_ = json.Unmarshal(byteValue, &c)

	return c
}

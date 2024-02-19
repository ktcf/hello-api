package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	Port            string `json:"port"`
	DefaultLanguage string `json:"default_language"`
	LegacyEndpoint  string `json:"legacy_endpoint"`
	DatabaseType    string `json:"database_type"`
	DatabaseURL     string `json:"database_url"`
}

var defaultConfiguration = Configuration{
	Port:            ":8080",
	DefaultLanguage: "english",
}

// LoadFromEnv will load configuration from the environment
func (c *Configuration) LoadFromEnv() {
	if lang := os.Getenv("DEFAULT_LANGUAGE"); lang != "" {
		c.DefaultLanguage = lang
	}

	if port := os.Getenv("PORT"); port != "" {
		c.Port = port
	}
}

// ParsePort will check to see if the port is in proper format and a number
func (c *Configuration) ParsePort() {
	if c.Port[0] != ':' {
		c.Port = ":" + c.Port
	}

	if _, err := strconv.Atoi(c.Port[1:]); err != nil {
		fmt.Printf("invalid port %s", c.Port)
	}
}

// LoadFromJSON will read a JSON file and update the configuration based on the file
func (c *Configuration) LoadFromJSON(path string) error {
	log.Printf("loading configuration from file: %s\n", path)

	b, err := os.ReadFile(path)
	if err != nil {
		log.Printf("unable to load file: %v\n", err)
		return errors.New("unable to load configuration")
	}

	if err = json.Unmarshal(b, c); err != nil {
		log.Printf("unable to parse file: %v\n", err)
		return errors.New("unable to load configuration")
	}

	// Verify require fields
	if c.Port == "" {
		log.Printf("empty port, reverting to default")
		c.Port = defaultConfiguration.Port
	}

	if c.DefaultLanguage == "" {
		log.Printf("empty language, reverting to default")
		c.DefaultLanguage = defaultConfiguration.DefaultLanguage
	}

	return nil
}

// LoadConfiguration will provide cycle through flags, files, and finally env variables
func LoadConfiguration() Configuration {
	cfgFileFlag := flag.String("config_file", "", "load configuration from a file")
	portFlag := flag.String("port", "", "set port")

	flag.Parse()

	cfg := defaultConfiguration

	if cfgFileFlag != nil && *cfgFileFlag != "" {
		if err := cfg.LoadFromJSON(*cfgFileFlag); err != nil {
			log.Printf("unable to load configuration from json: %s, using default values", *cfgFileFlag)
		}
	}

	cfg.LoadFromEnv()

	if portFlag != nil && *portFlag != "" {
		cfg.Port = *portFlag
	}

	cfg.ParsePort()
	return cfg
}

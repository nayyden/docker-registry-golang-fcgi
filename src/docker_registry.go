package main

import (
	"encoding/json"
	"net/http/fcgi"
	"os"
)

var logger = &Logger{}

type Config struct {
	DataDir  string
	LogLevel string
}

func readConfig() (*Config, error) {
	configuration := Config{}

	// Read config file
	file, err := os.Open("conf.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func startServer(cfg *Config) {
	logger.Info("Using dataDir ", cfg.DataDir)

	if err := fcgi.Serve(nil, NewHandler(cfg.DataDir)); err != nil {
		logger.Error(err.Error())
	}
}

func main() {
	// Read configuration file
	cfg, err := readConfig()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Set log level from config
	switch cfg.LogLevel {
	case "info":
		logger.Level = INFO
		break
	case "debug":
		logger.Level = DEBUG
		break
	case "error":
		logger.Level = ERROR
		break
	case "warn":
		logger.Level = WARN
		break
	}

	// Everything above this line is executed only on the request and reused
	// in the next requests due to the way fastcgi manages processes.
	startServer(cfg)
}

package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

const (
	configPath = "./configs/global.yaml"
)

type GCFG struct {
	API struct {
		ClientTimeout time.Duration
	}
	Server struct {
		Address     string `yaml:"addr"`
		Port        string `yaml:"port"`
		WebRoot     string `yaml:"root"`
		AppPrefix   string `yaml:"appPfx"`
		ApiPrefix   string `yaml:"apiPfx"`
		AdminPrefix string `yaml:"admPfx"`
	} `yaml:"server"`
}

func (cfg *GCFG) Init() error {
	cfgFile, err := os.Open(configPath)

	if errors.Is(err, os.ErrNotExist) {
		log.Print("Generating Default Configs")
		cfg.initDefaultConfig()
		return nil
	} else if err != nil {
		return err
	}
	defer cfgFile.Close()

	var configContent []byte
	_, err = cfgFile.Read(configContent)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configContent, cfg)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *GCFG) initDefaultConfig() {
	cfg.API.ClientTimeout = 500000

	cfg.Server.Address = ""
	cfg.Server.Port = "8080"
	cfg.Server.WebRoot = "./web/"
	cfg.Server.AppPrefix = "/app"
	cfg.Server.ApiPrefix = "/api"
	cfg.Server.AdminPrefix = "/admin"
}

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

type ServerConfig struct {
	Address     string `yaml:"addr"`
	Port        string `yaml:"port"`
	WebFileRoot string `yaml:"root"`
	AppPrefix   string `yaml:"appPfx"`
	ApiPrefix   string `yaml:"apiPfx"`
	AdminPrefix string `yaml:"admPfx"`
}

type GCFG struct {
	API struct {
		ClientTimeout time.Duration
	}
	*ServerConfig `yaml:"server"`
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

	cfg.ServerConfig = &ServerConfig{
		Address:     "",
		Port:        "8080",
		WebFileRoot: "./web/",
		AppPrefix:   "/app",
		ApiPrefix:   "/api",
		AdminPrefix: "/admin",
	}
}

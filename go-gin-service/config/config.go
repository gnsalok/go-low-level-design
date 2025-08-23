package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	JWT struct {
		SecretKey string `yaml:"secretKey"`
		Issuer    string `yaml:"issuer"`
	} `yaml:"jwt"`
}

func LoadConfig() (*Config, error) {
	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
		log.Println("config.yaml not found, creating a default one.")
		defaultConfig := `
server:
  port: ":8080"
jwt:
  secretKey: "a_very_secret_key_change_me"
  issuer: "my-app"
`
		if err := os.WriteFile("config.yaml", []byte(defaultConfig), 0644); err != nil {
			return nil, fmt.Errorf("failed to write default config file: %w", err)
		}
	}
	f, err := os.Open("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %w", err)
	}
	defer f.Close()
	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("could not decode config file: %w", err)
	}
	return &cfg, nil
}

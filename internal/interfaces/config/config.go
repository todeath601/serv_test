package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server *serverConfig `json:"server"`
}

type serverConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	ReadTimeout  int    `json:"readTimeout"` //time in seconds
	WriteTimeout int    `json:"writeTimeout"`
}

func LoadConfig() (Config, error) {

	c := &Config{
		Server: &serverConfig{},
	}

	data, err := os.ReadFile("/Users/nick/GoVolodya/internal/interfaces/config/config.json")
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return Config{}, err
	}
	return *c, nil
}

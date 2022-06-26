package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ProxyConfig struct {
	Routes map[string]Route
}

type Route struct {
	Name    string
	Port    int
	Address string
	Type    string
	Headers map[string]Header `yaml:"headers,omitempty"`
	Paths   []string          `yaml:"paths,omitempty"`
}

type Header struct {
	Name   string   `yaml:"name"`
	Values []string `yaml:"values"`
}

func LoadConfig() (*ProxyConfig, error) {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	config := &ProxyConfig{}
	err = yaml.Unmarshal(file, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

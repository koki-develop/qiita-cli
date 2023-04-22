package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AccessToken string `yaml:"access_token"`
	Format      string `yaml:"format"`
}

func Dir() (string, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(h, ".qiita-cli"), nil
}

func Path() (string, error) {
	d, err := Dir()
	if err != nil {
		return "", err
	}

	return path.Join(d, "config.yaml"), nil
}

func Load() (*Config, error) {
	p, err := Path()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

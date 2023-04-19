package config

import (
	"fmt"
	"os"
	"path"
	"syscall"

	"golang.org/x/term"
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

func Configure(cfg *Config) error {
	if cfg.AccessToken == "" {
		fmt.Print("Qiita Access Token:")
		passwd, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		fmt.Println()
		cfg.AccessToken = string(passwd)
	}
	if cfg.Format == "" {
		var f string
		fmt.Print("Default Output Format: ")
		if _, err := fmt.Scanln(&f); err != nil {
			return err
		}
		cfg.Format = f
	}

	y, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	d, err := Dir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(d, os.ModePerm); err != nil {
		return err
	}

	p, err := Path()
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(y); err != nil {
		return err
	}
	fmt.Printf("Configured! (%s)\n", p)

	return nil
}

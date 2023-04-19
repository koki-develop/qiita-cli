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

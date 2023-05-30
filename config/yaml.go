package config

import (
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Hostname string        `yaml:"hostname"`
	Port     int           `yaml:"port"`
	Timeout  time.Duration `yaml:"timeout"`
}

// DefaultConfigurationFile returns the default configuration file
func DefaultConfigurationFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".flsvc.yaml"), nil
}

// ValidateConfigurationFile validates the configuration file
func (c *Conf) ValidateConfigurationFile(targetConfigFile string) (*Conf, error) {
	buf, err := os.ReadFile(targetConfigFile)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(buf, c); err != nil {
		return nil, err
	}
	return c, nil
}

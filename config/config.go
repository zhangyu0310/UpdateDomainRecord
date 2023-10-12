package config

import (
	"encoding/json"
	"errors"
	"os"
	"sync/atomic"
)

var globalCfg atomic.Value

var (
	ErrAccessKeyIDEmpty     = errors.New("access key id is empty")
	ErrAccessKeySecretEmpty = errors.New("access key secret is empty")
	ErrDomainNameEmpty      = errors.New("domain name is empty")
)

type Config struct {
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	DomainName      string `json:"domain_name"`
	// Endpoint 请参考 https://api.aliyun.com/product/Alidns
	EndPoint  string `json:"end_point"`
	Frequency uint   `json:"frequency"`
	LogPath   string `json:"log_path"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Valid() error {
	if c.AccessKeyID == "" {
		return ErrAccessKeyIDEmpty
	}
	if c.AccessKeySecret == "" {
		return ErrAccessKeySecretEmpty
	}
	if c.DomainName == "" {
		return ErrDomainNameEmpty
	}
	if c.EndPoint == "" {
		c.EndPoint = "alidns.cn-hangzhou.aliyuncs.com"
	}
	if c.Frequency == 0 {
		c.Frequency = 1
	}
	if c.LogPath == "" {
		c.LogPath = "./"
	}
	return nil
}

// InitializeConfigFromFile read config from file.
func InitializeConfigFromFile(configFile string) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	cfg := NewConfig()
	if err = json.Unmarshal(data, cfg); err != nil {
		return err
	}

	if err = cfg.Valid(); err != nil {
		return err
	}

	StoreGlobalConfig(cfg)
	return nil
}

// GetGlobalConfig returns the global configuration for this server.
// It should store configuration from command line and configuration file.
// Other parts of the system can read the global configuration use this function.
func GetGlobalConfig() *Config {
	return globalCfg.Load().(*Config)
}

// StoreGlobalConfig stores a new config to the globalConf. It mostly uses in the test to avoid some data races.
func StoreGlobalConfig(config *Config) {
	globalCfg.Store(config)
}

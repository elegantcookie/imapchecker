package imapchecker

import (
	"crypto/tls"
	"time"
)

type Config struct {
	Address   string
	UseTLS    bool
	TLSConfig *tls.Config

	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

func NewConfig(address string) *Config {
	return &Config{
		Address:      address,
		UseTLS:       false,
		TLSConfig:    nil,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func NewTLSConfig(address string, tlsConfig *tls.Config) *Config {
	return &Config{
		Address:      address,
		UseTLS:       true,
		TLSConfig:    tlsConfig,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func parseConfig(c *Config) error {
	if c == nil {
		return NoConfig
	}
	if c.Address == "" {
		return EmptyAddress
	}
	if c.UseTLS {
		if c.TLSConfig == nil {
			return NoTLSConfig
		}
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = 15 * time.Second
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = 15 * time.Second
	}

	return nil
}

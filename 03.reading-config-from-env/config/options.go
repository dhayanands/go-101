package config

import (
	"os"
	"strings"
)

type Option func(cfg *Config) (string, error)

func WithDBConfig() Option {
	return func(cfg *Config) (string, error) {
		// store config
		s := os.Getenv(keyDBHost)
		if s == "" {
			return keyDBHost, ErrMissingEnvProp
		}
		cfg.dbHost = strings.TrimSpace(s)

		s = os.Getenv(keyDBPort)
		if s == "" {
			return keyDBPort, ErrMissingEnvProp
		}
		cfg.dbPort = strings.TrimSpace(s)

		s = os.Getenv(keyDBPass)
		if s == "" {
			return keyDBPass, ErrMissingEnvProp
		}
		cfg.dbPass = strings.TrimSpace(s)

		s = os.Getenv(keyDBUser)
		if s == "" {
			return keyDBUser, ErrMissingEnvProp
		}
		cfg.dbUser = strings.TrimSpace(s)

		s = os.Getenv(keyDBSchema)
		if s == "" {
			return keyDBSchema, ErrMissingEnvProp
		}
		cfg.dbSchema = strings.TrimSpace(s)

		return "", nil
	}
}

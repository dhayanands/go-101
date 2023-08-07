package config

import (
	"errors"
	"os"
)

var (
	ErrMissingEnvProp = errors.New("missing.env.prop")
)

const (
	keyDBHost   string = "DB_PG_HOST"
	keyDBPort   string = "DB_PG_PORT"
	keyDBUser   string = "DB_PG_USERNAME"
	keyDBPass   string = "DB_PG_PASSWORD"
	keyDBSchema string = "DB_PG_DATABASE"
	keyEnv      string = "ENV"
)

type Config struct {
	dbUser   string
	dbPass   string
	dbHost   string
	dbPort   string
	dbSchema string
	env      string
}

func (c *Config) DBUser() string   { return c.dbUser }
func (c *Config) DBPass() string   { return c.dbPass }
func (c *Config) DBHost() string   { return c.dbHost }
func (c *Config) DBPort() string   { return c.dbPort }
func (c *Config) DBSchema() string { return c.dbSchema }
func (c *Config) Env() string      { return c.env }

func New(opts ...Option) (*Config, string, error) {
	cfg := new(Config)
	cfg.env = os.Getenv(keyEnv)

	for _, opt := range opts {
		k, err := opt(cfg)
		if err != nil {
			return nil, k, err
		}
	}
	return cfg, "", nil
}

func main() {
	return
}

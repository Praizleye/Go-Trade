// Package config loads runtime configuration from environment variables
// (and from a .env file in the working directory, if present).
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all runtime settings for the bot.
// Fields are populated by Load() from environment variables.
type Config struct {
	BinanceAPIKey    string
	BinanceAPISecret string
	UseTestnet       bool
	DBPath           string
	LogLevel         string
}

// Load reads configuration from environment variables, falling back
// to defaults where sensible. It returns an error if any required
// value is missing or malformed.
func Load() (*Config, error) {
	// Try to load .env from the working directory. If the file doesn't
	// exist, that's not an error — env vars might be set directly
	// (e.g., in production by systemd or Docker).
	_ = godotenv.Load()

	cfg := &Config{
		BinanceAPIKey:    os.Getenv("BINANCE_API_KEY"),
		BinanceAPISecret: os.Getenv("BINANCE_API_SECRET"),
		UseTestnet:       getEnvBool("BINANCE_USE_TESTNET", true),
		DBPath:           getEnvOrDefault("DB_PATH", "./bot.db"),
		LogLevel:         getEnvOrDefault("LOG_LEVEL", "info"),
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Validate checks that required fields are present.
// Returns the first missing-field error it finds.
func (c *Config) Validate() error {
	if c.BinanceAPIKey == "" {
		return fmt.Errorf("BINANCE_API_KEY is required (set it in .env)")
	}
	if c.BinanceAPISecret == "" {
		return fmt.Errorf("BINANCE_API_SECRET is required (set it in .env)")
	}
	return nil
}

// getEnvOrDefault returns the env var if set and non-empty, else fallback.
func getEnvOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// getEnvBool parses an env var as a boolean ("true", "false", "1", "0").
// Falls back to the provided default if unset or unparseable.
func getEnvBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(v)
	if err != nil {
		return fallback
	}
	return parsed
}

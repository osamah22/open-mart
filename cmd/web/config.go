package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	DB_URL      string `env:"DB_URL,required"`
	SESSION_KEY string `env:"SESSION_KEY,required"`
	ENV         string `env:"ENV" envDefault:"development"`
	PORT        int    `env:"PORT" envDefault:"8080"`

	GOOGLE_CLIENT_ID     string `env:"GOOGLE_CLIENT_ID,required"`
	GOOGLE_CLIENT_SECRET string `env:"GOOGLE_CLIENT_SECRET,required"`
	BASE_URL             string `env:"BASE_URL" envDefault:"http://localhost:8080"`
}

func LoadConfig() (Config, error) {
	var cfg Config

	// ----------------------------
	// Flags (highest priority)
	// ----------------------------
	pflag.String("db-url", "", "Postgres connection URL")
	pflag.String("session-key", "", "Session encryption key")
	pflag.String("env", "development", "Environment")
	pflag.Int("port", 8080, "HTTP server port")

	pflag.String("google-client-id", "", "Google OAuth client ID")
	pflag.String("google-client-secret", "", "Google OAuth client secret")
	pflag.String("base-url", "http://localhost:8080", "Base URL of the app")
	pflag.Parse()

	// ----------------------------
	// Viper setup
	// ----------------------------
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".") // fallback

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.BindPFlags(pflag.CommandLine)

	// Read YAML (optional but expected)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	// ----------------------------
	// Assign values
	// ----------------------------
	cfg = Config{
		DB_URL:               viper.GetString("db-url"),
		SESSION_KEY:          viper.GetString("session-key"),
		ENV:                  viper.GetString("env"),
		PORT:                 viper.GetInt("port"),
		GOOGLE_CLIENT_ID:     viper.GetString("google-client-id"),
		GOOGLE_CLIENT_SECRET: viper.GetString("google-client-secret"),
		BASE_URL:             viper.GetString("base-url"),
	}

	// ----------------------------
	// Validation
	// ----------------------------
	if cfg.DB_URL == "" {
		return Config{}, fmt.Errorf("db-url is required")
	}
	if cfg.SESSION_KEY == "" {
		return Config{}, fmt.Errorf("session-key is required")
	}

	if cfg.GOOGLE_CLIENT_ID == "" {
		return Config{}, fmt.Errorf("google-client-id is required")
	}
	if cfg.GOOGLE_CLIENT_SECRET == "" {
		return Config{}, fmt.Errorf("google-client-secret is required")
	}
	return cfg, nil
}

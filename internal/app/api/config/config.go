package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"strings"
	"time"
)

type Config struct {
	PostgresDSN string
	GRPC        GRPCConfig
	/// naming как в Makefile
	MigrationsPath string
	AppSecretKey   string
	TokenTTL       time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GRPCConfig struct {
	Port int `yaml:"port"`
}

func MustLoad() *Config {
	config := &Config{
		PostgresDSN:    GetEnvString("POSTGRES_DSN"),
		GRPC:           GRPCConfig{},
		MigrationsPath: "",
		AppSecretKey:   "",
		TokenTTL:       0,
	}

	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

func GetEnvString(name string, defaultVal string) string {
	val := os.Getenv(name)
	if strings.TrimSpace(val) == "" {
		return defaultVal
	}

	return val
}

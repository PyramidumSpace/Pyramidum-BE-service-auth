package config

import (
        "fmt"
	"time"
        "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env            string 		`env:"ENV" 		env-default:"local"`
	GRPC           GRPCConfig
        Storage        StorageConfig
	StoragePath    string
	MigrationsPath string 		`env:"MIGRATIONS_PATH" 	env-default:"./migrations"`
	AppSecretKey   string 		`env:"APP_SECRET_KEY" 	env-default:"very-secret-key"`
	TokenTTL       time.Duration 	`env:"TOKEN_TTL" 	env-default:"1h"`
}

type StorageConfig struct {
       Dialect	string `env:"DB_DIALECT" 	env-default:"postgres"`
       User     string `env:"DB_USER" 		env-required:"true"`
       Password string `env:"DB_PASSWORD" 	env-required:"true"`
       Host     string `env:"DB_HOST" 		env-required:"true"`
       Port     string `env:"DB_PORT" 		env-default:"5432"`
       Database string `env:"DB_DATABASE" 	env-required:"true"`
       SSLMode  string `env:"DB_SSL" 		env-default:"disable"`
}

type GRPCConfig struct {
	Port    int `env:"GRPC_PORT" env-required:"true"`
}

func MustLoad() *Config {
	var cfg Config

        cleanenv.ReadEnv(&cfg)

        cfg.StoragePath = fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Storage.Dialect,
		cfg.Storage.User,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.Database,
		cfg.Storage.SSLMode,               
	)
        return &cfg
}

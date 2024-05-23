package config

import (
	"os"
	"time"
        "strconv"
)

type Config struct {
	Env            string
	StoragePath    string
	GRPC           GRPCConfig
	MigrationsPath string
	AppSecretKey   string
	TokenTTL       time.Duration
}

type GRPCConfig struct {
	Port    int
}

func MustLoad() *Config {
	var cfg Config

        cfg.Env = os.Getenv("ENV")
        cfg.StoragePath = os.Getenv("STORAGE_PATH")
        
        var grpcCfg GRPCConfig
        port, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
        if err != nil {
            panic(err)
        }

        grpcCfg.Port = port
        
        cfg.GRPC = grpcCfg
        cfg.MigrationsPath = os.Getenv("MIGRATIONS_PATH")
        cfg.AppSecretKey = os.Getenv("APP_SECRET_KEY")

        token_ttl, err := time.ParseDuration(os.Getenv("TOKEN_TTL"))
        if err != nil {
            panic(err)
        }
        cfg.TokenTTL = token_ttl
        return &cfg
}

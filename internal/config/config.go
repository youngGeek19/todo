package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"prod"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"4s"`
	Timeout     time.Duration `yaml:"timeout" env-default:"60s"`
}

func MustLoad() *Config {
	if err := godotenv.Load(".env.dev"); err != nil {
		log.Fatalf("Cannot load the env file! %s: ", err)
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is didn't set!")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesn't exist! %s: ", err)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Couldn't read the config file! %s: ", err)
	}

	return &cfg
}

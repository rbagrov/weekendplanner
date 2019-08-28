package settings

import (
	env "github.com/caarlos0/env"
	log "log"
)

// Config struct to handle all gin default config
type Config struct {
	DbUser     string `env:"DBUSER" envDefault:"postgres"`
	DbPassword string `env:"DBPASSWORD" envDefault:"postgres"`
	DbName     string `env:"DBName" envDefault:"weekendplanner"`
	DbHost     string `env:"DBHost" envDefault:"localhost"`
	DbPort     string `env:"DBPort" envDefault:"5432"`
}

// Settings function that set project settings
func Settings() Config {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Fatal("Settings could not be parsed")
	}
	return config

}

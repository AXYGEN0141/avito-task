package config

import (
	"os"
)

// Config is a configuration structure for DataBase.
type Config struct {
	Mode   string
	Port   string
	DBConf *DBConf
}

// DBConf is a structure for DataBase configuration.
type DBConf struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// GetConfig reads config from .env file.
// If failed gets default value.
func GetConfig() *Config {
	return &Config{
		Mode: getEnvAsStr("APP_MODE", "debug"),
		Port: getEnvAsStr("APP_PORT", ":6969"),
		DBConf: &DBConf{
			Dialect:  getEnvAsStr("POSTGRES_DIALECT", "pgx"),
			Host:     getEnvAsStr("POSTGRES_URI", "127.0.0.1"),
			Port:     getEnvAsStr("POSTGRES_PORT", "5432"),
			Username: getEnvAsStr("POSTGRES_USER", "postgres"),
			Password: getEnvAsStr("POSTGRES_PASSWORD", "postgres"),
			DBName:   getEnvAsStr("POSTGRES_DB", "testdb"),
		},
	}
}

// getEnvAsStr represents Environmental variables as String
func getEnvAsStr(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

package config

import "os"

type ConfigMySQL struct {
	MySqlUser     string
	MySqlPassword string
	MySqlHost     string
	MySqlPort     string
	MySqlProtocol string
	MySqlDB       string
	AppPort       string
}

func LoadConfigMySQL() *ConfigMySQL {
	return &ConfigMySQL{
		MySqlUser:     getEnv("MYSQL_USER", "root"),
		MySqlPassword: getEnv("MYSQL_PASS", "root"),
		MySqlHost:     getEnv("MYSQL_HOST", "localhost"),
		MySqlPort:     getEnv("MYSQL_PORT", "3306"),
		MySqlProtocol: getEnv("MYSQL_PROTOCOL", "tcp"),
		MySqlDB:       getEnv("MYSQL_DB", "userdb"),
		AppPort:       getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

package config

import (
	"log"
	"os"
	"strconv"
)

// GetEnv mengembalikan nilai ENV (development / production)
func GetEnv() string {
	return getEnvironmentValue("ENV")
}

// GetDataSourceURL mengembalikan database connection URL
func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

// GetApplicationPort mengembalikan port aplikasi (int)
func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}

// helper untuk ambil env, fail fast kalau tidak ada
func getEnvironmentValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return value
}


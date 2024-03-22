package utils

import "os"

// IsProductionEnv returns whether the app is running using production env
func IsProductionEnv() bool {
	return os.Getenv("APP_ENV") == "production"
}

// GetEnv returns app envorinment : e.g. development, production, staging, testing, etc
func GetEnv() string {
	return os.Getenv("APP_ENV")
}

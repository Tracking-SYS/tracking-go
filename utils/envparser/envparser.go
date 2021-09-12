package envparser

import (
	"os"
	"strconv"
	"strings"
)

//GetString ...
func GetString(key string, defaultValue string) string {
	envValue := os.Getenv(key)
	if strings.TrimSpace(envValue) == "" {
		return defaultValue
	}

	return envValue
}

//GetInt ...
func GetInt(key string, defaultValue int) int {
	envValue := os.Getenv(key)
	if strings.TrimSpace(envValue) == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(envValue)
	if err != nil {
		return defaultValue
	}
	return intValue
}

//GetFloat ...
func GetFloat(key string, defaultValue float64) float64 {
	envValue := os.Getenv(key)
	if strings.TrimSpace(envValue) == "" {
		return defaultValue
	}

	v, err := strconv.ParseFloat(envValue, 64)
	if err != nil {
		return defaultValue
	}

	return v
}

//GetArray ...
func GetArray(key string, separator string, defaultValue []string) []string {
	envValue := os.Getenv(key)
	if strings.TrimSpace(envValue) == "" {
		return defaultValue
	}

	return strings.Split(envValue, separator)
}

//GetBool ...
func GetBool(key string, defaultValue bool) bool {
	envValue := strings.TrimSpace(os.Getenv(key))
	if envValue == "" {
		return defaultValue
	}

	v, err := strconv.ParseBool(envValue)
	if err != nil {
		return defaultValue
	}

	return v
}

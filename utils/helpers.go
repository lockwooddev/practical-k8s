package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetEnv(key string, required bool) string {
	value := os.Getenv(key)
	if value == "" && required {
		logrus.Fatalf("env var '%s' is not set or empty", key)
	}
	return value
}

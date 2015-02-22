package asset

import (
	"os"
	"strings"
)

func readValue(flags int) (string, bool) {
	if flags&Environment != 0 {
		if value, ok := readFromEnvironment(); ok {
			return value, true
		}
	}
	return "", false
}

func readFromEnvironment() (string, bool) {
	return GetEnv(EnvironmentVariable)
}

func GetEnv(key string) (string, bool) {
	environment := os.Environ()
	for _, value := range environment {
		parts := strings.Split(value, "=")
		if len(parts) == 2 && parts[0] == key {
			return parts[1], true
		}
	}
	return "", false
}

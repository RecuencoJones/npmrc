package main

import "os"

func GetEnv(name, defaultValue string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return defaultValue
}

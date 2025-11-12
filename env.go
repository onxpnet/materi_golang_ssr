package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var envCache map[string]string

// loadEnvFile reads .env once and stores it in envCache.
func loadEnvFile(filename string) {
	envCache = make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		log.Println("No .env file found")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		envCache[key] = val
	}
}

// GetEnv checks OS env first, then .env cache, else default.
func GetEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	if val, ok := envCache[key]; ok {
		return val
	}
	return defaultVal
}

package app

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load()
}

func GetEnv(key string, def ...string) string {
	var d string
	if len(def) > 0 {
		d = def[0]
	}

	res := os.Getenv(key)

	if res == "" {
		return d
	}

	return res
}

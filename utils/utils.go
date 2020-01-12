package utils

import (
	"fmt"
	"os"
	"os/user"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func RelativeToHome(path string) string {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	return fmt.Sprintf("%s/%s", homeDir, path)
}

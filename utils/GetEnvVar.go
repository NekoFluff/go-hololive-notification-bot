package utils

import (
	"fmt"
	"log"
	"os"
)

func GetEnvVar(name string) string {
	envVar := os.Getenv(name)
	if envVar == "" {
		log.Fatal(fmt.Sprintf("$%v must be set", name))
	}
	return envVar
}
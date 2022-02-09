package environment

import (
	"fmt"
	"os"
)

const (
	Port    = "PORT"
	WebPort = "WEB_PORT"
)

func CheckEnvVars() {
	envVars := []string{
		Port,
		WebPort,
	}

	for _, v := range envVars {
		if os.Getenv(v) == "" {
			panic(fmt.Sprintf("env variable %s must be defined", v))
		}
	}
}

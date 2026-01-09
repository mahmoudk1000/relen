package utils

import (
	"fmt"
	"strings"
)

func ParseProjectSlashApplication(args []string) (string, string, error) {
	if len(args) == 0 {
		return "", "", fmt.Errorf("no arguments provided")
	}

	if len(args) == 1 {
		parsed := strings.Split(args[0], "/")

		return parsed[0], parsed[1], nil
	}

	return args[0], args[1], nil
}

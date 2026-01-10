package utils

import (
	"fmt"
	"strings"
)

func ParseProjectSlashApplication(args []string) (project, app string, err error) {
	if len(args) == 0 {
		return "", "", fmt.Errorf("no arguments provided")
	}

	if len(args) == 1 && strings.Contains(args[0], "/") {
		parts := strings.SplitN(args[0], "/", 2)
		if len(parts) != 2 {
			return "", "", fmt.Errorf("invalid format: use 'project/app' or 'project app'")
		}

		return parts[0], parts[1], nil
	}

	if len(args) == 2 {
		return args[0], args[1], nil
	}

	return "", "", fmt.Errorf("invalid format: use 'project/app' or 'project app'")
}

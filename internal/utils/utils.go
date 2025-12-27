package utils

import (
	"os"
	"path/filepath"
)

const (
	baseVerDBConfigDir   = "verdb"
	projectsFileName     = "projects.json"
	applicationsFileName = "applications.json"
)

func configDirBuilder(configFileName string) (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(configDir, baseVerDBConfigDir, configFileName)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return "", err
	}

	return path, nil
}

func ProjectConfigFilePath() (string, error) {
	return configDirBuilder(projectsFileName)
}

func ApplicationConfigFilePath() (string, error) {
	return configDirBuilder(applicationsFileName)
}

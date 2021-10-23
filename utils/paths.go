package utils

import (
	"os"
	"path"
	"path/filepath"
)

func GetRuntimeDir() (string, error) {
	ex, _ := os.Executable()
	abspath, err := filepath.Abs(path.Dir(ex))
	if err != nil {
		return "", err
	}
	return abspath, nil
}

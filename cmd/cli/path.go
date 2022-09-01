package cli

import (
	"os"
	"path/filepath"
)

const defaultAppFolder = "notify-go"

func GetConfigPath() string {
	path, _ := os.UserConfigDir()
	return filepath.Dir(path + "/" + defaultAppFolder)
}

func SuffixConfigPath(suffix string) string {
	return filepath.Dir(GetConfigPath() + "/" + suffix)
}

func InitFolderPath() error {
	path := GetConfigPath()
	return os.MkdirAll(path, os.ModeDir)
}

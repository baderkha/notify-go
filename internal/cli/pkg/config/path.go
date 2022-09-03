package config

import (
	"os"
	"path/filepath"
)

const defaultAppFolder = "notify-go"

func GetPath() string {
	path, _ := os.UserConfigDir()
	return filepath.Join(path, defaultAppFolder)

}

func InitFolderPath() error {
	path := GetPath()
	return os.MkdirAll(path, os.ModePerm)
}

package config

import (
	"fmt"
	"github.com/gohook/gohook/config/configfile"
	"io"
	"os"
	"path/filepath"
)

const (
	ConfigFileName = "gohook.json"
	configFileDir  = ".gohook"
)

var (
	configDir = os.Getenv("GOHOOK_CONFIG")
)

func init() {
	if configDir == "" {
		configDir = filepath.Join(GetHomeDir(), configFileDir)
	}
}

func GetHomeDir() string {
	return os.Getenv("HOME")
}

func ConfigDir() string {
	return configDir
}

func SetConfigDir(dir string) {
	configDir = dir
}

func NewConfigFile(filename string) *configfile.ConfigFile {
	return &configfile.ConfigFile{
		Host:     "gohook.io:9001",
		Filename: filename,
	}
}

func Load(configDir string) (*configfile.ConfigFile, error) {
	if configDir == "" {
		configDir = ConfigDir()
	}

	configFile := NewConfigFile(filepath.Join(configDir, ConfigFileName))
	_, err := os.Stat(configFile.Filename)
	if err != nil {
		if os.IsNotExist(err) {
			return configFile, fmt.Errorf("config file not found - %s", configFile.Filename)
		}

		return configFile, fmt.Errorf("%s - %v", configFile.Filename, err)
	}

	file, err := os.Open(configFile.Filename)
	if err != nil {
		return configFile, fmt.Errorf("%s - %v", configFile.Filename, err)
	}
	defer file.Close()

	err = configFile.LoadFromReader(file)
	if err != nil {
		if err == io.EOF {
			return configFile, fmt.Errorf("config file is empty - %s", configFile.Filename)
		}
		return configFile, fmt.Errorf("%s - %v", configFile.Filename, err)
	}

	return configFile, nil
}

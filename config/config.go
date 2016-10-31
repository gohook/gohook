package config

import (
	"fmt"
	"github.com/gohook/gohook/config/configfile"
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
		Filename: filename,
	}
}

func Load(configDir string) (*configfile.ConfigFile, error) {
	if configDir == "" {
		configDir = ConfigDir()
	}

	configFile := configfile.ConfigFile{
		Filename: filepath.Join(configDir, ConfigFileName),
	}

	_, err := os.Stat(configFile.Filename)
	if err == nil {
		file, err := os.Open(configFile.Filename)
		if err != nil {
			return &configFile, fmt.Errorf("%s - %v", configFile.Filename, err)
		}
		defer file.Close()
		err = configFile.LoadFromReader(file)
		if err != nil {
			err = fmt.Errorf("%s - %v", configFile.Filename, err)
		}
		return &configFile, err
	} else if !os.IsNotExist(err) {
		return &configFile, fmt.Errorf("%s - %v", configFile.Filename, err)
	}
	return &configFile, nil
}

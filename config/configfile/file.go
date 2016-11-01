package configfile

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ConfigFile struct {
	AuthToken string `json:"auth_token"`
	Filename  string `json:"-"`
	Host      string `json:"host"`
}

func (configFile *ConfigFile) LoadFromReader(configData io.Reader) error {
	return json.NewDecoder(configData).Decode(&configFile)
}

func (configFile *ConfigFile) ContainsAuthToken() bool {
	return configFile.AuthToken != ""
}

func (configFile *ConfigFile) SaveToWriter(writer io.Writer) error {
	data, err := json.Marshal(configFile)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	return err
}

func (configFile *ConfigFile) Save() error {
	if configFile.Filename == "" {
		return fmt.Errorf("Can't save config with out a filename")
	}

	if err := os.MkdirAll(filepath.Dir(configFile.Filename), 0700); err != nil {
		return err
	}

	f, err := os.OpenFile(configFile.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	return configFile.SaveToWriter(f)
}

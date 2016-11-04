package command

import (
	"fmt"
	cliflags "github.com/gohook/gohook/cli/flags"
	"github.com/gohook/gohook/client"
	"github.com/gohook/gohook/config"
	"github.com/gohook/gohook/config/configfile"
)

type GohookCli struct {
	configFile *configfile.ConfigFile
	client     *client.GohookClient
}

func (cli *GohookCli) Initialize(opts *cliflags.ClientOptions, store client.HookStore) error {
	var err error
	cli.configFile, err = LoadDefaultConfigFile()
	if err != nil {
		return err
	}

	client, err := client.NewGohookClient(cli.configFile, store)
	if err != nil {
		return err
	}
	cli.client = client
	return nil
}

func (cli *GohookCli) Client() *client.GohookClient {
	return cli.client
}

func (cli *GohookCli) ConfigFile() *configfile.ConfigFile {
	return cli.configFile
}

func NewGohookCli() *GohookCli {
	return &GohookCli{}
}

func LoadDefaultConfigFile() (*configfile.ConfigFile, error) {
	configFile, err := config.Load(config.ConfigDir())
	if err != nil {
		return nil, err
	}
	if !configFile.ContainsAuthToken() {
		return nil, fmt.Errorf("config file is missing token")
	}
	return configFile, nil
}

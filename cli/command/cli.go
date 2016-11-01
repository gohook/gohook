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
	cli.configFile = LoadDefaultConfigFile()
	client, err := client.NewGohookClient(cli.configFile.AuthToken, store)
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

func LoadDefaultConfigFile() *configfile.ConfigFile {
	configFile, err := config.Load(config.ConfigDir())
	if err != nil {
		fmt.Println("Error loading config file:", err)
	}
	if !configFile.ContainsAuthToken() {
		fmt.Println("Missing authentication token. Please log in")
	}
	return configFile
}

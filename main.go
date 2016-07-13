package main

import (
	"fmt"
	"github.com/gohook/gohook/commands"
	"github.com/urfave/cli"
	"os"
)

var (
	appName        = "gohook"
	appDescription = "Run commands from webhooks"

	// Version for applciation
	Version = "v0.0.1"
)

func EntryCommand(c *cli.Context) error {
	if c.Bool("start") {
		// Look for start flag to kickoff app
		return c.App.Command("start").Run(c)
	}
	return cli.ShowAppHelp(c)
}

func StartCommand(c *cli.Context) error {
	fmt.Println("Starting client")
	fmt.Println("Exiting because there is nothing here yet")
	return nil
}

func main() {
	app := cli.NewApp()

	// Application CLI Config
	app.Name = appName
	app.Version = Version
	app.Action = EntryCommand
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Brian Egizi",
			Email: "me@brianegizi.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "$HOME/.gohook/",
			Usage: "location of gohook.json file",
		},
		cli.BoolFlag{
			Name:  "start, s",
			Usage: "run the client",
		},
	}

	// Define application commands
	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Run the client",
			Action:  StartCommand,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a webhook command",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "output, o", Usage: "Define the output file for command's STDOUT"},
				cli.StringSliceFlag{Name: "params, p", Usage: "Define params to be collected from the webhook as passed into STDIN"},
			},
			Action: commands.Add,
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "Remove a webhook command",
			Action:  commands.Remove,
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List all webhook commands",
			Action:  commands.List,
		},
	}

	app.Run(os.Args)
}

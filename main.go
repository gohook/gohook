package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/gohook/gohook-server/client"
	"github.com/gohook/gohook/commands"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var (
	appName        = "gohook"
	appDescription = "Run commands from webhooks"

	// Version for applciation
	Version = "v0.0.1"
)

func main() {

	// Establish GRPC Connection
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	app := cli.NewApp()
	service := client.New(conn, log.NewNopLogger())

	// Application CLI Config
	app.Name = appName
	app.Version = Version
	app.Action = commands.EntryCommand
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
			Action:  commands.StartCommand(service),
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a webhook command",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "output, o", Usage: "Define the output file for command's STDOUT"},
				cli.StringSliceFlag{Name: "params, p", Usage: "Define params to be collected from the webhook as passed into STDIN"},
			},
			Action: commands.Add(service),
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
			Action:  commands.List(service),
		},
	}

	app.Run(os.Args)
}

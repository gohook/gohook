package commands

import (
	"github.com/urfave/cli"
)

func EntryCommand(c *cli.Context) error {
	if c.Bool("start") {
		// Look for start flag to kickoff app
		return c.App.Command("start").Run(c)
	}
	return cli.ShowAppHelp(c)
}

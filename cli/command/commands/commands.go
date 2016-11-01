package commands

import (
	"github.com/gohook/gohook/cli/command"
	"github.com/gohook/gohook/cli/command/add"
	"github.com/gohook/gohook/cli/command/list"
	"github.com/gohook/gohook/cli/command/start"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command, gohookCli *command.GohookCli) {
	cmd.AddCommand(
		list.NewListCommand(gohookCli),
		add.NewAddCommand(gohookCli),
		start.NewStartCommand(gohookCli),
	)
}

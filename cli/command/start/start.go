package start

import (
	"github.com/gohook/gohook/cli/command"
	"github.com/spf13/cobra"
)

type startOptions struct{}

func NewStartCommand(gohookCli *command.GohookCli) *cobra.Command {
	opts := &startOptions{}

	return &cobra.Command{
		Use:   "start",
		Short: "Run tunnel to get notified when webhooks get hit",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStart(gohookCli, opts)
		},
	}
}

func runStart(gohookCli *command.GohookCli, opts *startOptions) error {
	return gohookCli.Client().Tunnel()
}

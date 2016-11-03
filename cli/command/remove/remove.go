package remove

import (
	"errors"
	"fmt"
	"github.com/gohook/gohook/cli/command"
	"github.com/spf13/cobra"
)

type removeOptions struct {
	id string
}

func NewRemoveCommand(gohookCli *command.GohookCli) *cobra.Command {
	opts := &removeOptions{}

	return &cobra.Command{
		Use:   "remove [hookId]",
		Short: "Remove an existing hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("hookId is required")
			}
			opts.id = args[0]
			return runRemove(gohookCli, opts)
		},
	}
}

func runRemove(gohookCli *command.GohookCli, opts *removeOptions) error {
	err := gohookCli.Client().Remove(opts.id)
	if err != nil {
		return err
	}

	fmt.Printf("Removed hook %s\n", opts.id)
	return nil
}

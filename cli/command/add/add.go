package add

import (
	"fmt"
	"github.com/gohook/gohook/cli/command"
	"github.com/spf13/cobra"
)

type addOptions struct {
	method  string
	command []string
}

func NewAddCommand(gohookCli *command.GohookCli) *cobra.Command {
	opts := &addOptions{}

	cmd := &cobra.Command{
		Use:   "add [command]",
		Short: "Add a new command hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.command = args
			return runAdd(gohookCli, opts)
		},
	}

	flags := cmd.Flags()

	flags.StringVarP(&opts.method, "method", "m", "GET", "Set the http method this hooks listens for (defaults to GET)")

	return cmd
}

func runAdd(gohookCli *command.GohookCli, opts *addOptions) error {
	hook, err := gohookCli.Client().Create(opts.method, opts.command)
	if err != nil {
		return err
	}

	// formatter.Write() <- write out to stdout
	fmt.Printf("Add command, ID: %s, Method: %s, Command: %s\n", hook.Id, hook.Method, hook.Command)
	return nil
}

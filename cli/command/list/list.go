package list

import (
	"fmt"
	"github.com/gohook/gohook/cli/command"
	"github.com/gohook/gohook/client"
	"github.com/spf13/cobra"
	// "golang.org/x/net/context"
)

type listOptions struct {
	all bool
}

func NewListCommand(gohookCli *command.GohookCli) *cobra.Command {
	opts := &listOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List hooks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(gohookCli, opts)
		},
	}

	flags := cmd.Flags()

	flags.BoolVarP(&opts.all, "all", "a", false, "Show all hooks (default hides hooks not set on this system)")

	return cmd
}

func runList(gohookCli *command.GohookCli, opts *listOptions) error {
	hooks, err := gohookCli.Client().List()
	if err != nil {
		return err
	}
	return formatHookTable(hooks)
}

func formatHookTable(hooks client.HookList) error {
	// This is lame. Upgrate to github.com/olekukonko/tablewriter in the future
	table := ""
	for i, h := range hooks {
		table += fmt.Sprintf("%d | %s | %s | %s |\n", i+1, h.Id, h.Method, h.Url)
	}

	fmt.Println(table)
	return nil
}

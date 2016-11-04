package main

import (
	"fmt"
	"github.com/gohook/gohook/cli/command"
	"github.com/gohook/gohook/cli/command/commands"
	cliflags "github.com/gohook/gohook/cli/flags"
	"github.com/gohook/gohook/filestore"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	gohookCli := command.NewGohookCli()
	cmd := newGohookCommand(gohookCli)

	err := cmd.Execute()
	if err != nil {
		fmt.Println("Command Error: ", err)
		os.Exit(1)
	}
}

func newGohookCommand(gohookCli *command.GohookCli) *cobra.Command {
	opts := cliflags.NewClientOptions()

	cmd := &cobra.Command{
		Use:           "gohook [OPTIONS] COMMAND [arg...]",
		Short:         "Create webhooks to run things.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Version {
				fmt.Println("Show Version")
				return nil
			}
			fmt.Println(cmd.UsageString())
			return nil
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			store, err := filestore.NewLocalHookStore(opts.HookDir)
			if err != nil {
				return err
			}

			return gohookCli.Initialize(opts, store)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.Version, "version", "v", false, "Print version information and quit")
	flags.StringVar(&opts.HookDir, "hook-dir", opts.HookDir, "Set location where hooks get stored")

	cmd.SetFlagErrorFunc(FlagErrorFunc)
	cmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	cmd.PersistentFlags().MarkShorthandDeprecated("help", "please use --help")

	commands.AddCommands(cmd, gohookCli)

	return cmd
}

func FlagErrorFunc(cmd *cobra.Command, err error) error {
	if err == nil {
		return err
	}

	usage := ""
	if cmd.HasSubCommands() {
		usage = "\n\n" + cmd.UsageString()
	}

	fmt.Printf("%s\nSee '%s --help'.%s\n", err, cmd.CommandPath(), usage)
	return nil
}

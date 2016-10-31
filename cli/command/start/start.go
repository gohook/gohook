package start

import (
	"fmt"
	"github.com/gohook/gohook/cli/command"
	"github.com/spf13/cobra"
	"io"
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
	stream, err := gohookCli.Client().Tunnel()
	if err != nil {
		return err
	}

	// Block and run forever
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println("Error", err)
			return err
		}
		hookCall := response.GetHook()
		fmt.Printf("Got Hit: %s, Method: %s\n", hookCall.Id, hookCall.Method)
	}

	return nil
}

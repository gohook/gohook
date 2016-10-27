package commands

import (
	"context"
	"fmt"
	"io"

	"github.com/gohook/gohook-server/client"
	"github.com/gohook/gohook-server/pb"
	"github.com/urfave/cli"
)

func StartCommand(s client.GohookClient) cli.ActionFunc {
	return func(c *cli.Context) error {
		fmt.Println("Starting client")
		stream, err := s.Tunnel(context.Background(), &pb.TunnelRequest{"myid"})
		if err != nil {
			fmt.Println(err)
			return err
		}

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
			fmt.Println(response.Event)
		}

		return nil
	}
}

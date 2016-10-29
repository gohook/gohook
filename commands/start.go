package commands

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gohook/gohook-server/client"
	"github.com/gohook/gohook-server/pb"
	"github.com/urfave/cli"
)

func StartCommand(s client.GohookClient) cli.ActionFunc {
	return func(c *cli.Context) error {
		fmt.Println("Starting client")
		stream, err := s.Tunnel(context.Background(), &pb.TunnelRequest{})
		if err != nil {
			fmt.Println(err)
			return err
		}

		var rawCommand string
		if len(os.Args) > 1 {
			rawCommand = strings.Join(c.Args(), " ")
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

			cmd := exec.Command("sh", "-c", rawCommand)
			cmd.Stdout = os.Stdout

			err = cmd.Run()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(response.Event)
		}

		return nil
	}
}

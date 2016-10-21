package commands

import (
	"fmt"

	"github.com/gohook/gohook-server/gohookd"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
)

func List(s gohookd.Service) cli.ActionFunc {
	return func(c *cli.Context) error {
		fmt.Println("Listing all webhook commands")

		list, err := s.List(context.Background())
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}

		fmt.Println("Res: ", list)

		return nil
	}
}

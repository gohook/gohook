package commands

import (
	"fmt"
	"github.com/urfave/cli"
)

func Remove(c *cli.Context) error {
	fmt.Println("remove webhook command: ", c.Args().First())
	return nil
}

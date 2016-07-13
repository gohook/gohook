package commands

import (
	"fmt"
	"github.com/urfave/cli"
)

func List(c *cli.Context) error {
	fmt.Println("Listing all webhook commands")
	return nil
}

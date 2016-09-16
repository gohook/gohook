package commands

import (
	"fmt"
	"github.com/urfave/cli"
)

func StartCommand(c *cli.Context) error {
	fmt.Println("Starting client")
	fmt.Println("Exiting because there is nothing here yet")
	return nil
}

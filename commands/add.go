package commands

import (
	"fmt"
	"github.com/urfave/cli"
)

func Add(c *cli.Context) error {
	fmt.Println("args: ", c.Args())
	fmt.Println("flags: ", c.FlagNames())
	return nil
}

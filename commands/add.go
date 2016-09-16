package commands

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Add(c *cli.Context) error {
	fmt.Println("args: ", c.Args())
	fmt.Println("flags: ", c.FlagNames())

	var rawCommand string
	if len(os.Args) > 1 {
		rawCommand = strings.Join(c.Args(), " ")
	}

	cmd := exec.Command("sh", "-c", rawCommand)
	cmd.Stdout = os.Stdout

	fmt.Println("Running command in 60 seconds: ", rawCommand)

	time.Sleep(time.Minute * 1)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done.")
	return nil
}

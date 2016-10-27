package commands

import (
	"fmt"
	// "log"
	// "os"
	// "os/exec"
	// "strings"
	// "time"

	"github.com/gohook/gohook-server/client"
	"github.com/gohook/gohook-server/gohookd"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
)

func Add(s client.GohookClient) cli.ActionFunc {
	return func(c *cli.Context) error {
		fmt.Println("args: ", c.Args())
		fmt.Println("flags: ", c.FlagNames())

		newHook := gohookd.HookRequest{
			Method: "GET",
		}

		hook, err := s.Create(context.Background(), newHook)
		if err != nil {
			fmt.Println("Error: ", err)
			return err
		}

		fmt.Println("Res: ", hook)

		// var rawCommand string
		// if len(os.Args) > 1 {
		//	rawCommand = strings.Join(c.Args(), " ")
		// }

		// cmd := exec.Command("sh", "-c", rawCommand)
		// cmd.Stdout = os.Stdout

		// fmt.Println("Running command in 60 seconds: ", rawCommand)

		// time.Sleep(time.Minute * 1)

		// err := cmd.Run()
		// if err != nil {
		//	log.Fatal(err)
		// }

		// fmt.Println("done.")
		return nil
	}
}

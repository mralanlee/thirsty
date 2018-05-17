package main

import (
	"fmt"
	"log"
	"os"
	v "thirsty/validate"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "thirsty"
	app.Action = func(c *cli.Context) error {
		result := v.CheckEnv(c.Args().Get(0))

		fmt.Println(result)
		return nil
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}

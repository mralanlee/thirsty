package main

import (
	"fmt"
	"log"
	"os"
	"thirsty/validate"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "thirsty"
	app.Action = func(c *cli.Context) error {
		result := validate.CheckEnv(c.Args().Get(0))

		fmt.Println(result)
		return nil
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}

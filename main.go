package main

import (
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "thirsty"
	app.Action = func(c *cli.Context) error {

	}
}

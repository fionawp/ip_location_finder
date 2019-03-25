package main

import (
	"github.com/ip_location_finder/commands"
	"github.com/urfave/cli"
	"os"
)

var version = "development"

func main() {
	app := cli.NewApp()
	app.Name = "ip_location_finder"
	app.Usage = ""
	app.Version = version
	app.Copyright = "(c) 2019 The ip_location_finder contributors <fionawp@126.com>"
	app.EnableBashCompletion = true
	app.Flags = commands.GlobalFlags

	app.Commands = []cli.Command{
		commands.ConfigCommand,
		commands.StartCommand,
	}

	app.Run(os.Args)
}

package commands

import (
	"fmt"
	"github.com/ip_location_finder/context"
	"github.com/urfave/cli"
)

// Prints current configuration
var ConfigCommand = cli.Command{
	Name:   "config",
	Usage:  "Displays global configuration values",
	Action: configAction,
}

func configAction(ctx *cli.Context) error {
	conf := context.NewConfig(ctx)

	fmt.Printf("NAME                  VALUE\n")
	fmt.Printf("debug                 %t\n", conf.Debug())
	fmt.Printf("config-file           %s\n", conf.ConfigFile())
	fmt.Printf("app-name              %s\n", conf.AppName())
	fmt.Printf("app-version           %s\n", conf.AppVersion())
	fmt.Printf("app-copyright         %s\n", conf.AppCopyright())

	fmt.Printf("http-host             %s\n", conf.HttpServerHost())
	fmt.Printf("http-port             %d\n", conf.HttpServerPort())
	fmt.Printf("http-mode             %s\n", conf.HttpServerMode())

	return nil
}

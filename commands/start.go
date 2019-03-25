package commands

import (
	"fmt"
	"github.com/ip_location_finder/context"
	"github.com/ip_location_finder/server"
	"github.com/ip_location_finder/service"
	"github.com/urfave/cli"
	"log"
)

// Starts web server (user interface)
var StartCommand = cli.Command{
	Name:   "start",
	Usage:  "Starts web server",
	Flags:  startFlags,
	Action: startAction,
}

var startFlags = []cli.Flag{
	cli.IntFlag{
		Name:   "http-port, p",
		Usage:  "HTTP server port",
		Value:  8081,
		EnvVar: "IP_LOCATION_FINDER_HTTP_PORT",
	},
	cli.StringFlag{
		Name:   "http-host, i",
		Usage:  "HTTP server host",
		Value:  "",
		EnvVar: "IP_LOCATION_FINDER_HTTP_HOST",
	},
	cli.StringFlag{
		Name:   "http-mode, m",
		Usage:  "debug, release or test",
		Value:  "",
		EnvVar: "IP_LOCATION_FINDER_HTTP_MODE",
	},
}

func startAction(ctx *cli.Context) error {
	conf := context.NewConfig(ctx)

	//load ip file into memory
	fileName := conf.GetIpFile()
	myLogger := conf.GetLog()
	myLogger.Info("file  " + fileName)
	ipFileToList,_ := service.NewIpFileToList(fileName, conf)

	if conf.HttpServerPort() < 1 {
		log.Fatal("Server port must be a positive integer")
	}

	fmt.Printf("Starting web server at %s:%d...\n", ctx.String("http-host"), ctx.Int("http-port"))

	server.Start(conf, ipFileToList)

	fmt.Println("Done.")

	return nil
}

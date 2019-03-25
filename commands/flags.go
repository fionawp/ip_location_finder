package commands

import "github.com/urfave/cli"

// Global CLI flags
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "debug",
		Usage:  "run in debug mode",
		EnvVar: "IP_LOCATION_FINDER_DEBUG",
	},
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
	cli.StringFlag{
		Name:   "ip-file",
		Usage:  "ip file",
		Value:  "resources/ip2zipcode.txt",
		EnvVar: "IP_File",
	},
}

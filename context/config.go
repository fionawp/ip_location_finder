package context

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/urfave/cli"
	"os"
	"time"
)

// Config provides a struct in which application configuration is stored.
// Application code must use functions to get config values, for two reasons:
//
// 1. Some values are computed and we don't want to leak implementation details (aims at reducing refactoring overhead).
//
// 2. Paths might actually be dynamic later (if we build a multi-user version).
//
// See https://github.com/photoprism/photoprism/issues/50#issuecomment-433856358
type Config struct {
	appName        string
	appVersion     string
	appCopyright   string
	debug          bool
	configFile     string
	httpServerHost string
	httpServerPort int
	httpServerMode string
	ipFile         string
	myLog          *logging.Logger
}

// NewConfig() creates a new configuration entity by using two methods:
//
// 1. SetValuesFromFile: This will initialize values from a yaml config file.
//
// 2. SetValuesFromCliContext: Which comes after SetValuesFromFile and overrides
//    any previous values giving an option two override file configs through the CLI.
func NewConfig(ctx *cli.Context) *Config {
	c := &Config{}
	c.appName = ctx.App.Name
	c.appCopyright = ctx.App.Copyright
	c.appVersion = ctx.App.Version
	c.SetValuesFromCliContext(ctx)
	c.setLog()

	return c
}

// SetValuesFromCliContext uses values from the CLI to setup configuration overrides
// for the entity.
func (c *Config) SetValuesFromCliContext(ctx *cli.Context) error {
	if ctx.GlobalBool("debug") {
		c.debug = ctx.GlobalBool("debug")
	}

	if ctx.GlobalIsSet("http-host") || c.httpServerHost == "" {
		c.httpServerHost = ctx.GlobalString("http-host")
	}

	if ctx.GlobalIsSet("http-port") || c.httpServerPort == 0 {
		c.httpServerPort = ctx.GlobalInt("http-port")
	}

	if ctx.GlobalIsSet("http-mode") || c.httpServerMode == "" {
		c.httpServerMode = ctx.GlobalString("http-mode")
	}

	if ctx.GlobalIsSet("ip-file") || c.ipFile == "" {
		c.ipFile = ctx.GlobalString("ip-file")
	}

	return nil
}

// AppName returns the application name.
func (c *Config) AppName() string {
	return c.appName
}

// AppVersion returns the application version.
func (c *Config) AppVersion() string {
	return c.appVersion
}

// AppCopyright returns the application copyright.
func (c *Config) AppCopyright() string {
	return c.appCopyright
}

// Debug returns true if debug mode is on.
func (c *Config) Debug() bool {
	return c.debug
}

// ConfigFile returns the config file name.
func (c *Config) ConfigFile() string {
	return c.configFile
}

// HttpServerHost returns the built-in HTTP server host name or IP address (empty for all interfaces).
func (c *Config) HttpServerHost() string {
	return c.httpServerHost
}

// HttpServerPort returns the built-in HTTP server port.
func (c *Config) HttpServerPort() int {
	return c.httpServerPort
}

// HttpServerMode returns the server mode.
func (c *Config) HttpServerMode() string {
	return c.httpServerMode
}

func (c *Config) GetIpFile() string {
	return c.ipFile
}

type Log struct {
	Logger *logging.Logger
}

func (c *Config) setLog() {
	logFile := c.LogFilePath()
	backend1 := logging.NewLogBackend(logFile, "", 0)
	var format1 = logging.MustStringFormatter(
		`%{level:.4s} %{time:2006-01-02T15:04:05.999}  %{id:03x} %{message}`,
	)
	backend1Leveled := logging.NewBackendFormatter(backend1, format1)

	if c.debug {
		backend2 := logging.NewLogBackend(os.Stderr, "", 0)
		var format2 = logging.MustStringFormatter(
			`%{color}%{time:2006-01-02T15:04:05.999} %{level:.4s} %{id:03x}%{color:reset} %{message}`,
		)
		backend2Formatter := logging.NewBackendFormatter(backend2, format2)
		logging.SetBackend(backend2Formatter, backend1Leveled)
	} else {
		logging.SetBackend(backend1Leveled)
	}

	c.myLog = logging.MustGetLogger("example")
}

func (c *Config) GetLog() *logging.Logger {
	return c.myLog
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (c *Config) LogFilePath() *os.File {
	currentDate := time.Now().Format("2006-01-02")
	path := "logs/"
	isExist, _ := pathExists(path)
	if isExist == false {
		dirErr := os.MkdirAll(path, 0777)
		if nil != dirErr {
			fmt.Println(dirErr)
		}
	}
	fileName := path + "/" + currentDate + ".log"
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("fail to open log file " + fileName)
	}
	return logFile
}

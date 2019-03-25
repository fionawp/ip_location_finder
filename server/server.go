package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ip_location_finder/context"
	"github.com/ip_location_finder/service"
	"io"
)

// Start the REST API server using the configuration provided
func Start(conf *context.Config, file *service.IpFileToList) {
	if conf.HttpServerMode() != "" {
		gin.SetMode(conf.HttpServerMode())
	} else if conf.Debug() == false {
		gin.SetMode(gin.ReleaseMode)
	}

	logFile := conf.LogFilePath()
	gin.DefaultWriter = io.MultiWriter(logFile)
	app := gin.Default()

	registerRoutes(app, conf, file)

	app.Run(fmt.Sprintf("%s:%d", conf.HttpServerHost(), conf.HttpServerPort()))
}

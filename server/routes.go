package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ip_location_finder/apis"
	"github.com/ip_location_finder/context"
	"github.com/ip_location_finder/service"
)

func registerRoutes(app *gin.Engine, conf *context.Config, file *service.IpFileToList) {
	//routes
	searchPrefix := app.Group("/search")
	{
		apis.Location(searchPrefix, conf, file)
	}
}

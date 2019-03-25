package common

import (
	"github.com/gin-gonic/gin"
)

type result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type resultWithoutDate struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func FormatResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, &result{
		code,
		msg,
		data,
	})
}

func FormatResponseWithoutData(c *gin.Context, code int, msg string) {
	c.JSON(200, &resultWithoutDate{
		code,
		msg,
	})
	return
}

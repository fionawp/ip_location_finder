package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/ip_location_finder/common"
	"github.com/ip_location_finder/context"
	"github.com/ip_location_finder/service"
	"math"
	"math/big"
	"net"
)

type item struct {
	ZipCode  int
	Location string
}

func Location(router *gin.RouterGroup, conf *context.Config, file *service.IpFileToList) {
	router.GET("/location", func(c *gin.Context) {
		var myLogger = conf.GetLog()
		ipLocationMap := file.FileInfo
		ip := c.Query("ip")

		location,zipCode := binarySearch(ip, ipLocationMap, c)
		myLogger.Info(location)

		if location == "" || zipCode == 0 {
			common.FormatResponseWithoutData(c, 10001, "Not Found! ")
		} else{
			common.FormatResponse(c, 10000, "", &item{
				ZipCode:  zipCode,
				Location: location,
			})
		}
	})
}

func binarySearch(ip string, ipLocationMap []service.IpLocation, c *gin.Context) (string, int) {
	length := len(ipLocationMap)
	midIndex := int(math.Floor(float64(length) / 2))

	//IP TRANSFER TO INT 
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	intIP := ret.Int64()

	if intIP < ipLocationMap[0].GetFIpInt() || intIP > ipLocationMap[length-1].GetCIpInt() {
		return "", 0
	}

	if intIP >= ipLocationMap[midIndex].GetFIpInt() && intIP <= ipLocationMap[midIndex].GetCIpInt() {
		return ipLocationMap[midIndex].GetLocation(), ipLocationMap[midIndex].GetZipCode()
	} else if intIP < ipLocationMap[midIndex].GetFIpInt() {
		return  binarySearch(ip, ipLocationMap[0:midIndex], c)
	} else {
		return  binarySearch(ip, ipLocationMap[midIndex:], c)
	}
}

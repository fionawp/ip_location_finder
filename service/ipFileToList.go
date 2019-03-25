package service

import (
	"bufio"
	"github.com/ip_location_finder/context"
	"io"
	"math/big"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
)

type IpFileToList struct {
	FileInfo []IpLocation
}

type IpList []IpLocation

func (ipList IpList) Len() int {
	return len(ipList)
}
func (ipList IpList) Swap(i, j int) {
	ipList[i], ipList[j] = ipList[j], ipList[i]
}

func (ipList IpList) Less(i, j int) bool {
	return ipList[i].GetFIpInt() < ipList[j].GetFIpInt()
}

func NewIpFileToList(fileName string, conf *context.Config) (*IpFileToList, error) {
	var ipList IpList
	f, err := os.Open(fileName)
	myLogger := conf.GetLog()
	if err != nil {
		myLogger.Error(err.Error())
		return nil, err
	}
	buf := bufio.NewReader(f)

	for {
		line, err1 := buf.ReadString('\n')
		if err1 != nil {
			if err1 == io.EOF {
				break
			}
			myLogger.Error(err1.Error())
			break
		}
		line = strings.TrimSpace(line)
		ipLocationSlice := strings.Split(line, "\t")

		if ipLocationSlice == nil || len(ipLocationSlice) < 3 {
			myLogger.Error("Line Recognition error: " + line)
			break
		}

		tmp := strings.Split(ipLocationSlice[0], "-")
		fip := tmp[0]
		cip := tmp[1]
		FIpInt := IpToInt(fip)
		CIpInt := IpToInt(cip)
		ZipCode, _ := strconv.Atoi(ipLocationSlice[2])

		var mapValue IpLocation
		mapValue.SetCIpInt(CIpInt)
		mapValue.SetFIpInt(FIpInt)
		mapValue.SetLocation(ipLocationSlice[1])
		mapValue.SetZipCode(ZipCode)
		ipList = append(ipList, mapValue)
	}

	sort.Stable(ipList)

	instance := &IpFileToList{
		FileInfo: ipList,
	}
	return instance, nil
}

func IpToInt(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

package service

type IpLocation struct {
	fIpInt   int64
	cIpInt   int64
	location string
	zipCode  int
}

func (ipLocation *IpLocation) GetFIpInt() int64 {
	return ipLocation.fIpInt
}

func (ipLocation *IpLocation) GetCIpInt() int64 {
	return ipLocation.cIpInt
}

func (ipLocation *IpLocation) GetLocation() string {
	return ipLocation.location
}

func (ipLocation *IpLocation) GetZipCode() int {
	return ipLocation.zipCode
}

func (ipLocation *IpLocation) SetFIpInt(ipInt int64) {
	ipLocation.fIpInt = ipInt
}

func (ipLocation *IpLocation) SetCIpInt(ipInt int64) {
	ipLocation.cIpInt = ipInt
}

func (ipLocation *IpLocation) SetLocation(location string) {
	ipLocation.location = location
}

func (ipLocation *IpLocation) SetZipCode(zip int) {
	ipLocation.zipCode = zip
}

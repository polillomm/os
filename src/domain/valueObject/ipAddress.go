package valueObject

import (
	"errors"
	"net"

	voHelper "github.com/goinfinite/os/src/domain/valueObject/helper"
)

type IpAddress string

func NewIpAddress(value interface{}) (ipAddress IpAddress, err error) {
	stringValue, err := voHelper.InterfaceToString(value)
	if err != nil {
		return ipAddress, errors.New("IpAddressValueMustBeString")
	}

	parsedIpAddress := net.ParseIP(stringValue)
	if parsedIpAddress == nil {
		return ipAddress, errors.New("InvalidIpAddress")
	}

	return IpAddress(stringValue), nil
}

func NewLocalhostIpAddress() IpAddress {
	return IpAddress("127.0.0.1")
}

func (vo IpAddress) String() string {
	return string(vo)
}

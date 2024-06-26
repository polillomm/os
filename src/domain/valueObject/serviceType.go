package valueObject

import (
	"errors"
	"slices"
	"strings"
)

type ServiceType string

var ValidServiceTypes = []string{
	"application",
	"runtime",
	"database",
	"webserver",
	"mom",
	"monitoring",
	"logging",
	"security",
	"backup",
	"system",
	"other",
}

func NewServiceType(value string) (ServiceType, error) {
	value = strings.ToLower(value)
	if !slices.Contains(ValidServiceTypes, value) {
		return "", errors.New("InvalidServiceType")
	}
	return ServiceType(value), nil
}

func NewServiceTypePanic(value string) ServiceType {
	st, err := NewServiceType(value)
	if err != nil {
		panic(err)
	}
	return st
}

func (st ServiceType) String() string {
	return string(st)
}

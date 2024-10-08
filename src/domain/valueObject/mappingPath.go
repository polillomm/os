package valueObject

import (
	"errors"
	"regexp"
	"strings"

	voHelper "github.com/goinfinite/os/src/domain/valueObject/helper"
)

const mappingPathRegex string = `^[^\s<>;'":#{}?\[\]]{1,512}$`

type MappingPath string

func NewMappingPath(value interface{}) (mappingPath MappingPath, err error) {
	stringValue, err := voHelper.InterfaceToString(value)
	if err != nil {
		return mappingPath, errors.New("MappingPathMustBeString")
	}

	hasLeadingSlash := strings.HasPrefix(stringValue, "/")
	if !hasLeadingSlash {
		stringValue = "/" + stringValue
	}

	re := regexp.MustCompile(mappingPathRegex)
	if !re.MatchString(stringValue) {
		return mappingPath, errors.New("InvalidMappingPath")
	}

	return MappingPath(stringValue), nil
}

func (vo MappingPath) String() string {
	return string(vo)
}

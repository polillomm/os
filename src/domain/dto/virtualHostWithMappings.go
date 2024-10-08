package dto

import (
	"github.com/goinfinite/os/src/domain/entity"
)

type VirtualHostWithMappings struct {
	entity.VirtualHost
	Mappings []entity.Mapping `json:"mappings"`
}

func NewVirtualHostWithMappings(
	vhost entity.VirtualHost,
	mappings []entity.Mapping,
) VirtualHostWithMappings {
	return VirtualHostWithMappings{
		VirtualHost: vhost,
		Mappings:    mappings,
	}
}

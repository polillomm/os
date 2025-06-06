package vhostInfra

import (
	"testing"

	testHelpers "github.com/goinfinite/os/src/devUtils"
	"github.com/goinfinite/os/src/domain/dto"
	"github.com/goinfinite/os/src/domain/valueObject"
	infraHelper "github.com/goinfinite/os/src/infra/helper"
	internalDbInfra "github.com/goinfinite/os/src/infra/internalDatabase"
)

func TestVirtualHostCmdRepo(t *testing.T) {
	testHelpers.LoadEnvVars()

	persistentDbSvc, _ := internalDbInfra.NewPersistentDatabaseService()
	vhostCmdRepo := NewVirtualHostCmdRepo(persistentDbSvc)
	vhostQueryRepo := NewVirtualHostQueryRepo(persistentDbSvc)

	vhostName, _ := infraHelper.ReadPrimaryVirtualHostHostname()

	t.Run("Create", func(t *testing.T) {
		vhostType, _ := valueObject.NewVirtualHostType("top-level")
		operatorAccountId, _ := valueObject.NewAccountId(0)
		ipAddress := valueObject.IpAddressSystem

		err := vhostCmdRepo.Create(dto.NewCreateVirtualHost(
			vhostName, vhostType, nil, nil, operatorAccountId, ipAddress,
		))
		if err != nil {
			t.Errorf("ExpectingNoErrorButGot: %v", err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		vhostReadResponse, err := vhostQueryRepo.Read(dto.ReadVirtualHostsRequest{
			Pagination: dto.PaginationUnpaginated,
		})
		if err != nil || len(vhostReadResponse.VirtualHosts) == 0 {
			t.Errorf("ExpectingNoErrorButGot: %v", err)
		}

		isWildcard := true
		err = vhostCmdRepo.Update(dto.UpdateVirtualHost{
			Hostname:   vhostReadResponse.VirtualHosts[0].Hostname,
			IsWildcard: &isWildcard,
		})
		if err != nil {
			t.Errorf("ExpectingNoErrorButGot: %v", err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		vhostReadResponse, err := vhostQueryRepo.Read(dto.ReadVirtualHostsRequest{
			Pagination: dto.PaginationUnpaginated,
		})
		if err != nil || len(vhostReadResponse.VirtualHosts) == 0 {
			t.Errorf("ExpectingNoErrorButGot: %v", err)
		}

		err = vhostCmdRepo.Delete(vhostReadResponse.VirtualHosts[0].Hostname)
		if err != nil {
			t.Errorf("ExpectingNoErrorButGot: %v", err)
		}
	})
}

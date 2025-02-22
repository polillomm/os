package useCase

import (
	"errors"
	"log/slog"

	"github.com/goinfinite/os/src/domain/dto"
	"github.com/goinfinite/os/src/domain/repository"
)

func DeleteService(
	servicesQueryRepo repository.ServicesQueryRepo,
	servicesCmdRepo repository.ServicesCmdRepo,
	mappingCmdRepo repository.MappingCmdRepo,
	activityRecordCmdRepo repository.ActivityRecordCmdRepo,
	deleteDto dto.DeleteService,
) error {
	readFirstInstalledRequestDto := dto.ReadFirstInstalledServiceItemsRequest{
		ServiceName: &deleteDto.Name,
	}
	serviceEntity, err := servicesQueryRepo.ReadFirstInstalledItem(
		readFirstInstalledRequestDto,
	)
	if err != nil {
		return err
	}

	isSystemService := serviceEntity.Type.String() == "system"
	if isSystemService {
		return errors.New("SystemServicesCannotBeUninstalled")
	}

	err = mappingCmdRepo.DeleteAuto(deleteDto.Name)
	if err != nil {
		slog.Error("DeleteAutoMappingError", slog.Any("error", err))
		return errors.New("DeleteAutoMappingsInfraError")
	}

	err = servicesCmdRepo.Delete(deleteDto.Name)
	if err != nil {
		slog.Error("DeleteServiceError", slog.Any("error", err))
		return errors.New("DeleteServiceInfraError")
	}

	NewCreateSecurityActivityRecord(activityRecordCmdRepo).DeleteService(deleteDto)

	return nil
}

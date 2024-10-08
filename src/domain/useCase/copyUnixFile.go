package useCase

import (
	"errors"
	"log"

	"github.com/goinfinite/os/src/domain/dto"
	"github.com/goinfinite/os/src/domain/repository"
)

func CopyUnixFile(
	filesQueryRepo repository.FilesQueryRepo,
	filesCmdRepo repository.FilesCmdRepo,
	copyUnixFile dto.CopyUnixFile,
) error {
	err := filesCmdRepo.Copy(copyUnixFile)
	if err != nil {
		log.Printf("CopyUnixFileInfraError: %s", err.Error())
		return errors.New("CopyUnixFileInfraError")
	}

	fileSourcePath := copyUnixFile.SourcePath
	fileDestinationPath := copyUnixFile.DestinationPath
	log.Printf(
		"File '%s' (%s) copy created to '%s' with name '%s'.",
		fileSourcePath.GetFileName().String(),
		fileSourcePath.GetFileDir().String(),
		fileDestinationPath.GetFileName().String(),
		fileDestinationPath.GetFileDir().String(),
	)

	return nil
}

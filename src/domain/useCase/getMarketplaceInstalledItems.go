package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/entity"
	"github.com/speedianet/os/src/domain/repository"
)

func GetMarketplaceInstalledItems(
	marketplaceQueryRepo repository.MarketplaceQueryRepo,
) ([]entity.MarketplaceInstalledItem, error) {
	installedItems, err := marketplaceQueryRepo.GetInstalledItems()
	if err != nil {
		log.Printf("GetMarketplaceInstalledItemsError: %s", err.Error())
		return nil, errors.New("GetMarketplaceInstalledItemsInfraError")
	}

	return installedItems, nil
}
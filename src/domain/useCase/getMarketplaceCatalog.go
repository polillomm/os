package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/entity"
	"github.com/speedianet/os/src/domain/repository"
)

func GetMarketplaceCatalog(
	marketplaceQueryRepo repository.MarketplaceQueryRepo,
) ([]entity.MarketplaceCatalogItem, error) {
	catalogItems, err := marketplaceQueryRepo.GetCatalogItems()
	if err != nil {
		log.Printf("GetMarketplaceCatalogItemsError: %s", err.Error())
		return nil, errors.New("GetMarketplaceCatalogItemsInfraError")
	}

	return catalogItems, nil
}
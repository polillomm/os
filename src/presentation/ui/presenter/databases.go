package presenter

import (
	"log/slog"
	"net/http"

	"github.com/goinfinite/os/src/domain/dto"
	"github.com/goinfinite/os/src/domain/entity"
	"github.com/goinfinite/os/src/domain/valueObject"
	internalDbInfra "github.com/goinfinite/os/src/infra/internalDatabase"
	"github.com/goinfinite/os/src/presentation/service"
	uiHelper "github.com/goinfinite/os/src/presentation/ui/helper"
	"github.com/goinfinite/os/src/presentation/ui/page"
	presenterDto "github.com/goinfinite/os/src/presentation/ui/presenter/dto"
	"github.com/labstack/echo/v4"
)

type DatabasesPresenter struct {
	databaseService *service.DatabaseService
}

func NewDatabasesPresenter(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
	trailDbSvc *internalDbInfra.TrailDatabaseService,
) *DatabasesPresenter {
	return &DatabasesPresenter{
		databaseService: service.NewDatabaseService(persistentDbSvc, trailDbSvc),
	}
}

func (presenter *DatabasesPresenter) databaseOverviewFactory(
	rawDatabaseType string,
) (databaseOverview presenterDto.DatabaseOverview, err error) {
	databaseType, err := valueObject.NewDatabaseType(rawDatabaseType)
	if err != nil {
		return databaseOverview, err
	}

	isInstalled := false
	databaseEntities := []entity.Database{}
	databaseOverview = presenterDto.NewDatabaseOverview(
		databaseType, isInstalled, databaseEntities,
	)

	requestBody := map[string]interface{}{
		"dbType":       databaseType.String(),
		"itemsPerPage": 1000,
	}
	responseOutput := presenter.databaseService.Read(requestBody)
	if responseOutput.Status != service.Success {
		return databaseOverview, err
	}

	responseDto, assertOk := responseOutput.Body.(dto.ReadDatabasesResponse)
	if assertOk {
		databaseOverview.IsInstalled = true
		databaseOverview.Databases = responseDto.Databases
	}

	return databaseOverview, nil
}

func (presenter *DatabasesPresenter) Handler(c echo.Context) error {
	rawDatabaseType := "mariadb"
	if c.QueryParam("dbType") != "" {
		rawDatabaseType = c.QueryParam("dbType")
	}

	selectedDatabaseOverview, err := presenter.databaseOverviewFactory(
		rawDatabaseType,
	)
	if err != nil {
		slog.Error("DatabaseOverviewFactoryError", slog.String("err", err.Error()))
		return nil
	}

	pageContent := page.DatabasesIndex(selectedDatabaseOverview)
	return uiHelper.Render(c, pageContent, http.StatusOK)
}

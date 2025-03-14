package apiController

import (
	"errors"
	"log/slog"
	"net/http"

	internalDbInfra "github.com/goinfinite/os/src/infra/internalDatabase"
	apiHelper "github.com/goinfinite/os/src/presentation/api/helper"
	"github.com/goinfinite/os/src/presentation/service"
	"github.com/labstack/echo/v4"
)

type DatabaseController struct {
	persistentDbService *internalDbInfra.PersistentDatabaseService
	dbService           *service.DatabaseService
}

func NewDatabaseController(
	persistentDbService *internalDbInfra.PersistentDatabaseService,
	trailDbSvc *internalDbInfra.TrailDatabaseService,
) *DatabaseController {
	return &DatabaseController{
		persistentDbService: persistentDbService,
		dbService: service.NewDatabaseService(
			persistentDbService, trailDbSvc,
		),
	}
}

// GetDatabases godoc
// @Summary      GetDatabases
// @Description  List databases names, users and sizes.
// @Tags         database
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Param        dbType path string true "DatabaseType (like mysql, postgres)"
// @Success      200 {array} entity.Database
// @Router       /v1/database/{dbType}/ [get]
func (controller *DatabaseController) Read(c echo.Context) error {
	requestInputData, err := apiHelper.ReadRequestInputData(c)
	if err != nil {
		return err
	}

	return apiHelper.ServiceResponseWrapper(
		c, controller.dbService.Read(requestInputData),
	)
}

// CreateDatabase godoc
// @Summary      CreateDatabase
// @Description  Create a new database.
// @Tags         database
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        dbType path string true "DatabaseType (like mysql, postgres)"
// @Param        createDatabaseDto body dto.CreateDatabase true "All props are required."
// @Success      201 {object} object{} "DatabaseCreated"
// @Router       /v1/database/{dbType}/ [post]
func (controller *DatabaseController) Create(c echo.Context) error {
	requestInputData, err := apiHelper.ReadRequestInputData(c)
	if err != nil {
		return err
	}

	return apiHelper.ServiceResponseWrapper(
		c, controller.dbService.Create(requestInputData),
	)
}

// DeleteDatabase godoc
// @Summary      DeleteDatabase
// @Description  Delete a database.
// @Tags         database
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        dbType path string true "DatabaseType (like mysql, postgres)"
// @Param        dbName path string true "DatabaseName"
// @Success      200 {object} object{} "DatabaseDeleted"
// @Router       /v1/database/{dbType}/{dbName}/ [delete]
func (controller *DatabaseController) Delete(c echo.Context) error {
	requestInputData, err := apiHelper.ReadRequestInputData(c)
	if err != nil {
		return err
	}

	return apiHelper.ServiceResponseWrapper(
		c, controller.dbService.Delete(requestInputData),
	)
}

func (controller *DatabaseController) parseUserPrivileges(rawPrivileges interface{}) (
	rawPrivilegesStrSlice []string, err error,
) {
	rawUniquePrivilegeStr, assertOk := rawPrivileges.(string)
	if assertOk {
		return []string{rawUniquePrivilegeStr}, nil
	}

	rawPrivilegesStrSlice, assertOk = rawPrivileges.([]string)
	if assertOk {
		return rawPrivilegesStrSlice, nil
	}

	rawPrivilegesInterfaceSlice, assertOk := rawPrivileges.([]interface{})
	if !assertOk {
		return rawPrivilegesStrSlice, errors.New("PrivilegesMustBeStringOrStringSlice")
	}
	for _, rawPrivilege := range rawPrivilegesInterfaceSlice {
		rawPrivilegeStr, assertOk := rawPrivilege.(string)
		if !assertOk {
			slog.Debug("InvalidPrivilegeType", slog.Any("privilege", rawPrivilege))
			continue
		}
		rawPrivilegesStrSlice = append(rawPrivilegesStrSlice, rawPrivilegeStr)
	}

	return rawPrivilegesStrSlice, nil
}

// CreateDatabaseUser godoc
// @Summary      CreateDatabaseUser
// @Description  Create a new database user.
// @Tags         database
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        dbType path string true "DatabaseType (like mysql, postgres)"
// @Param        dbName path string true "DatabaseName"
// @Param        createDatabaseUserDto body dto.CreateDatabaseUser true "privileges is optional. When not provided, privileges will be 'ALL'."
// @Success      201 {object} object{} "DatabaseUserCreated"
// @Router       /v1/database/{dbType}/{dbName}/user/ [post]
func (controller *DatabaseController) CreateUser(c echo.Context) error {
	requestInputData, err := apiHelper.ReadRequestInputData(c)
	if err != nil {
		return err
	}

	rawPrivilegesSlice := []string{}
	if requestInputData["privileges"] != nil {
		rawPrivilegesSlice, err = controller.parseUserPrivileges(
			requestInputData["privileges"],
		)
		if err != nil {
			return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err.Error())
		}
	}
	requestInputData["privileges"] = rawPrivilegesSlice

	return apiHelper.ServiceResponseWrapper(
		c, controller.dbService.CreateUser(requestInputData),
	)
}

// DeleteDatabaseUser godoc
// @Summary      DeleteDatabaseUser
// @Description  Delete a database user.
// @Tags         database
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        dbType path string true "DatabaseType (like mysql, postgres)"
// @Param        dbName path string true "DatabaseName"
// @Param        dbUser path string true "DatabaseUsername to delete."
// @Success      200 {object} object{} "DatabaseUserDeleted"
// @Router       /v1/database/{dbType}/{dbName}/user/{dbUser}/ [delete]
func (controller *DatabaseController) DeleteUser(c echo.Context) error {
	requestInputData, err := apiHelper.ReadRequestInputData(c)
	if err != nil {
		return err
	}

	return apiHelper.ServiceResponseWrapper(
		c, controller.dbService.DeleteUser(requestInputData),
	)
}

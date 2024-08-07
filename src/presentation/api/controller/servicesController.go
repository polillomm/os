package apiController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/useCase"
	"github.com/speedianet/os/src/domain/valueObject"
	voHelper "github.com/speedianet/os/src/domain/valueObject/helper"
	internalDbInfra "github.com/speedianet/os/src/infra/internalDatabase"
	servicesInfra "github.com/speedianet/os/src/infra/services"
	vhostInfra "github.com/speedianet/os/src/infra/vhost"
	mappingInfra "github.com/speedianet/os/src/infra/vhost/mapping"
	apiHelper "github.com/speedianet/os/src/presentation/api/helper"
)

type ServicesController struct {
	persistentDbSvc *internalDbInfra.PersistentDatabaseService
}

func NewServicesController(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
) *ServicesController {
	return &ServicesController{
		persistentDbSvc: persistentDbSvc,
	}
}

// ReadServices	 godoc
// @Summary      ReadServices
// @Description  List installed services and their status.
// @Tags         services
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Success      200 {array} dto.InstalledServiceWithMetrics
// @Router       /v1/services/ [get]
func (controller *ServicesController) Read(c echo.Context) error {
	servicesQueryRepo := servicesInfra.NewServicesQueryRepo(controller.persistentDbSvc)
	servicesList, err := useCase.ReadServicesWithMetrics(servicesQueryRepo)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, servicesList)
}

// ReadInstallableServices	 godoc
// @Summary      ReadInstallableServices
// @Description  List installable services.
// @Tags         services
// @Security     Bearer
// @Accept       json
// @Produce      json
// @Success      200 {array} entity.InstallableService
// @Router       /v1/services/installables/ [get]
func (controller *ServicesController) ReadInstallables(c echo.Context) error {
	servicesQueryRepo := servicesInfra.NewServicesQueryRepo(controller.persistentDbSvc)
	servicesList, err := useCase.ReadInstallableServices(servicesQueryRepo)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, servicesList)
}

func parsePortBindings(bindings []interface{}) []valueObject.PortBinding {
	var svcPortBindings []valueObject.PortBinding
	for _, portBinding := range bindings {
		portBindingMap, assertOk := portBinding.(map[string]interface{})
		if !assertOk {
			panic("InvalidPortBinding")
		}

		port, err := valueObject.NewNetworkPort(portBindingMap["port"])
		if err != nil {
			panic(err)
		}
		portBindingStr := port.String()

		if portBindingMap["protocol"] != nil {
			protocol, err := valueObject.NewNetworkProtocol(portBindingMap["protocol"])
			if err != nil {
				panic(err)
			}

			portBindingStr += "/" + protocol.String()
		}

		svcPortBinding, err := valueObject.NewPortBinding(portBindingStr)
		if err != nil {
			panic(err)
		}
		svcPortBindings = append(svcPortBindings, svcPortBinding)
	}

	return svcPortBindings
}

// CreateInstallableService godoc
// @Summary      CreateInstallableService
// @Description  Install a new installable service.
// @Tags         services
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        createInstallableServiceDto	body dto.CreateInstallableService	true	"Only name is required.<br />If version is not provided, it will be 'lts'.<br />If portBindings is not provided, it wil be default service port bindings.<br />If autoCreateMapping is not provided, it will be 'true'."
// @Success      201 {object} object{} "InstallableServiceCreated"
// @Router       /v1/services/installables/ [post]
func (controller *ServicesController) CreateInstallable(c echo.Context) error {
	requiredParams := []string{"name"}
	requestBody, _ := apiHelper.ReadRequestBody(c)

	apiHelper.CheckMissingParams(requestBody, requiredParams)

	serviceName, err := valueObject.NewServiceName(requestBody["name"])
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
	}

	var svcVersionPtr *valueObject.ServiceVersion
	if requestBody["version"] != nil {
		version, err := valueObject.NewServiceVersion(requestBody["version"])
		if err != nil {
			return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
		}
		svcVersionPtr = &version
	}

	var svcStartupFilePtr *valueObject.UnixFilePath
	if requestBody["startupFile"] != nil {
		startupFile, err := valueObject.NewUnixFilePath(requestBody["startupFile"])
		if err != nil {
			return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
		}
		svcStartupFilePtr = &startupFile
	}

	var svcPortBindings []valueObject.PortBinding
	if requestBody["portBindings"] != nil {
		svcPortBindings = parsePortBindings(
			requestBody["portBindings"].([]interface{}),
		)
	}

	autoCreateMapping := true
	if requestBody["autoCreateMapping"] != nil {
		var err error
		autoCreateMapping, err = voHelper.InterfaceToBool(
			requestBody["autoCreateMapping"],
		)
		if err != nil {
			return apiHelper.ResponseWrapper(
				c, http.StatusBadRequest, "InvalidAutoCreateMapping",
			)
		}
	}

	createInstallableServiceDto := dto.NewCreateInstallableService(
		serviceName, []valueObject.ServiceEnv{}, svcPortBindings, svcVersionPtr,
		svcStartupFilePtr, nil, nil, nil, nil, &autoCreateMapping,
	)

	servicesQueryRepo := servicesInfra.NewServicesQueryRepo(controller.persistentDbSvc)
	servicesCmdRepo := servicesInfra.NewServicesCmdRepo(controller.persistentDbSvc)
	mappingQueryRepo := mappingInfra.NewMappingQueryRepo(controller.persistentDbSvc)
	mappingCmdRepo := mappingInfra.NewMappingCmdRepo(controller.persistentDbSvc)
	vhostQueryRepo := vhostInfra.NewVirtualHostQueryRepo(controller.persistentDbSvc)

	err = useCase.CreateInstallableService(
		servicesQueryRepo,
		servicesCmdRepo,
		mappingQueryRepo,
		mappingCmdRepo,
		vhostQueryRepo,
		createInstallableServiceDto,
	)

	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusCreated, "InstallableServiceCreated")
}

// CreateCustomService godoc
// @Summary      CreateCustomService
// @Description  Install a new custom service.
// @Tags         services
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        createCustomServiceDto	body dto.CreateCustomService	true	"name, type and startCmd is required.<br />If version is not provided, it will be 'lts'.<br />If portBindings is not provided, it wil be default service port bindings.<br />If autoCreateMapping is not provided, it will be 'true'."
// @Success      201 {object} object{} "CustomServiceCreated"
// @Router       /v1/services/custom/ [post]
func (controller *ServicesController) CreateCustom(c echo.Context) error {
	requiredParams := []string{"name", "type", "startCmd"}
	requestBody, _ := apiHelper.ReadRequestBody(c)

	apiHelper.CheckMissingParams(requestBody, requiredParams)

	serviceName, err := valueObject.NewServiceName(requestBody["name"].(string))
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
	}

	svcType, err := valueObject.NewServiceType(requestBody["type"])
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
	}

	startCmd, err := valueObject.NewUnixCommand(requestBody["startCmd"])
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
	}

	var svcVersionPtr *valueObject.ServiceVersion
	if requestBody["version"] != nil {
		version, err := valueObject.NewServiceVersion(requestBody["version"])
		if err != nil {
			return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
		}
		svcVersionPtr = &version
	}

	var svcPortBindings []valueObject.PortBinding
	if requestBody["portBindings"] != nil {
		svcPortBindings = parsePortBindings(
			requestBody["portBindings"].([]interface{}),
		)
	}

	autoCreateMapping := true
	if requestBody["autoCreateMapping"] != nil {
		var err error
		autoCreateMapping, err = voHelper.InterfaceToBool(
			requestBody["autoCreateMapping"],
		)
		if err != nil {
			return apiHelper.ResponseWrapper(
				c, http.StatusBadRequest, "InvalidAutoCreateMapping",
			)
		}
	}

	createCustomServiceDto := dto.NewCreateCustomService(
		serviceName, svcType, startCmd, []valueObject.ServiceEnv{}, svcPortBindings,
		nil, nil, nil, nil, nil, svcVersionPtr, nil, nil, nil, nil, nil, nil, nil, nil,
		&autoCreateMapping,
	)

	servicesQueryRepo := servicesInfra.NewServicesQueryRepo(controller.persistentDbSvc)
	servicesCmdRepo := servicesInfra.NewServicesCmdRepo(controller.persistentDbSvc)
	mappingQueryRepo := mappingInfra.NewMappingQueryRepo(controller.persistentDbSvc)
	mappingCmdRepo := mappingInfra.NewMappingCmdRepo(controller.persistentDbSvc)
	vhostQueryRepo := vhostInfra.NewVirtualHostQueryRepo(controller.persistentDbSvc)

	err = useCase.CreateCustomService(
		servicesQueryRepo,
		servicesCmdRepo,
		mappingQueryRepo,
		mappingCmdRepo,
		vhostQueryRepo,
		createCustomServiceDto,
	)

	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusCreated, "CustomServiceCreated")
}

// UpdateService godoc
// @Summary      UpdateService
// @Description  Update service details.
// @Tags         services
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        updateServiceDto	body dto.UpdateService	true	"Only name is required.<br />Solo services can only change status.<br />status may be 'running', 'stopped' or 'uninstalled'."
// @Success      200 {object} object{} "ServiceUpdated"
// @Router       /v1/services/ [put]
func (controller *ServicesController) Update(c echo.Context) error {
	requiredParams := []string{"name"}
	requestBody, _ := apiHelper.ReadRequestBody(c)

	apiHelper.CheckMissingParams(requestBody, requiredParams)

	svcName, err := valueObject.NewServiceName(requestBody["name"])
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
	}

	var svcTypePtr *valueObject.ServiceType
	if requestBody["type"] != nil {
		svcType, err := valueObject.NewServiceType(requestBody["type"])
		if err != nil {
			return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
		}
		svcTypePtr = &svcType
	}

	var startCmdPtr *valueObject.UnixCommand
	if requestBody["startCmd"] != nil {
		startCmd, err := valueObject.NewUnixCommand(requestBody["startCmd"])
		if err != nil {
			return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err.Error())
		}
		startCmdPtr = &startCmd
	}

	var svcStatusPtr *valueObject.ServiceStatus
	if requestBody["status"] != nil {
		svcStatus := valueObject.NewServiceStatusPanic(
			requestBody["status"].(string),
		)
		svcStatusPtr = &svcStatus
	}

	var svcVersionPtr *valueObject.ServiceVersion
	if requestBody["version"] != nil {
		version, err := valueObject.NewServiceVersion(requestBody["version"])
		if err != nil {
			return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
		}
		svcVersionPtr = &version
	}

	var svcPortBindings []valueObject.PortBinding
	if requestBody["portBindings"] != nil {
		svcPortBindings = parsePortBindings(
			requestBody["portBindings"].([]interface{}),
		)
	}

	var startupFilePtr *valueObject.UnixFilePath
	if requestBody["startupFile"] != nil {
		startupFile, err := valueObject.NewUnixFilePath(requestBody["startupFile"])
		if err != nil {
			return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
		}
		startupFilePtr = &startupFile
	}

	updateSvcDto := dto.NewUpdateService(
		svcName, svcTypePtr, svcVersionPtr, svcStatusPtr, startCmdPtr, []valueObject.ServiceEnv{},
		svcPortBindings, nil, nil, nil, nil, nil, nil, nil, startupFilePtr, nil, nil,
		nil, nil, nil, nil,
	)

	servicesQueryRepo := servicesInfra.NewServicesQueryRepo(controller.persistentDbSvc)
	servicesCmdRepo := servicesInfra.NewServicesCmdRepo(controller.persistentDbSvc)
	mappingQueryRepo := mappingInfra.NewMappingQueryRepo(controller.persistentDbSvc)
	mappingCmdRepo := mappingInfra.NewMappingCmdRepo(controller.persistentDbSvc)

	err = useCase.UpdateService(
		servicesQueryRepo,
		servicesCmdRepo,
		mappingQueryRepo,
		mappingCmdRepo,
		updateSvcDto,
	)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, "ServiceUpdated")
}

// DeleteService godoc
// @Summary      DeleteService
// @Description  Delete/Uninstall a service.
// @Tags         services
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        svcName path string true "ServiceName to delete"
// @Success      200 {object} object{} "ServiceDeleted"
// @Router       /v1/services/{svcName}/ [delete]
func (controller *ServicesController) Delete(c echo.Context) error {
	svcName, err := valueObject.NewServiceName(c.Param("svcName"))
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusBadRequest, err)
	}

	servicesQueryRepo := servicesInfra.NewServicesQueryRepo(controller.persistentDbSvc)
	servicesCmdRepo := servicesInfra.NewServicesCmdRepo(controller.persistentDbSvc)
	mappingCmdRepo := mappingInfra.NewMappingCmdRepo(controller.persistentDbSvc)

	err = useCase.DeleteService(
		servicesQueryRepo,
		servicesCmdRepo,
		mappingCmdRepo,
		svcName,
	)
	if err != nil {
		return apiHelper.ResponseWrapper(c, http.StatusInternalServerError, err.Error())
	}

	return apiHelper.ResponseWrapper(c, http.StatusOK, "ServiceDeleted")
}

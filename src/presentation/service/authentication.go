package service

import (
	"github.com/goinfinite/os/src/domain/dto"
	"github.com/goinfinite/os/src/domain/useCase"
	"github.com/goinfinite/os/src/domain/valueObject"
	accountInfra "github.com/goinfinite/os/src/infra/account"
	activityRecordInfra "github.com/goinfinite/os/src/infra/activityRecord"
	authInfra "github.com/goinfinite/os/src/infra/auth"
	internalDbInfra "github.com/goinfinite/os/src/infra/internalDatabase"
	serviceHelper "github.com/goinfinite/os/src/presentation/service/helper"
)

type AuthService struct {
	trailDbSvc *internalDbInfra.TrailDatabaseService
}

func NewAuthService(
	trailDbSvc *internalDbInfra.TrailDatabaseService,
) *AuthService {
	return &AuthService{
		trailDbSvc: trailDbSvc,
	}
}

func (service *AuthService) GenerateJwtWithCredentials(
	input map[string]interface{},
) ServiceOutput {
	requiredParams := []string{"username", "password", "ipAddress"}
	err := serviceHelper.RequiredParamsInspector(input, requiredParams)
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	username, err := valueObject.NewUsername(input["username"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	password, err := valueObject.NewPassword(input["password"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	ipAddress, err := valueObject.NewIpAddress(input["ipAddress"])
	if err != nil {
		return NewServiceOutput(UserError, err.Error())
	}

	dto := dto.NewLogin(username, password, ipAddress)

	authQueryRepo := authInfra.AuthQueryRepo{}
	authCmdRepo := authInfra.AuthCmdRepo{}
	accQueryRepo := accountInfra.AccQueryRepo{}
	activityRecordQueryRepo := activityRecordInfra.NewActivityRecordQueryRepo(
		service.trailDbSvc,
	)
	activityRecordCmdRepo := activityRecordInfra.NewActivityRecordCmdRepo(
		service.trailDbSvc,
	)

	accessToken, err := useCase.GetSessionToken(
		authQueryRepo, authCmdRepo, accQueryRepo, activityRecordQueryRepo,
		activityRecordCmdRepo, dto,
	)
	if err != nil {
		return NewServiceOutput(InfraError, err.Error())
	}

	return NewServiceOutput(Success, accessToken)
}

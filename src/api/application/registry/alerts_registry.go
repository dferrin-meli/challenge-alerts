package registry

import (
	"challenge/alerts/src/api/alerts/api"
	"challenge/alerts/src/api/alerts/repository"
	"challenge/alerts/src/api/alerts/service"
	"challenge/alerts/src/api/application/common"
	"challenge/alerts/src/api/application/conf"
)

func CreateAlertsHandler(cfg *conf.Data) api.AlertsHandler {
	validate := common.CreateValidator()
	dbClient := common.NewDBClient(cfg.ConfigurationDB)
	alertRepository := repository.NewAlertsRepository(cfg, dbClient)
	alertService := service.NewAlertsService(alertRepository)

	return api.NewAlertsHandler(validate, alertService)
}

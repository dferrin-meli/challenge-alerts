package registry

import (
	"challenge/alerts/src/api/alerts/api"
	"challenge/alerts/src/api/application/common"
	"challenge/alerts/src/api/application/conf"
)

func CreateAlertsHandler(cfg *conf.Data) api.AlertsHandler {
	validate := common.CreateValidator()

	return api.NewAlertsHandler(validate)
}

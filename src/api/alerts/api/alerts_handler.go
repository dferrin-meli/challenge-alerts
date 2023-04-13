package api

import (
	"challenge/alerts/src/api/alerts/domain"
	"challenge/alerts/src/api/application/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AlertsHandler struct {
	alertService domain.AlertsService
	validate     *validator.Validate
}

func NewAlertsHandler(validate *validator.Validate, alertService domain.AlertsService) AlertsHandler {
	return AlertsHandler{
		alertService: alertService,
		validate:     validate,
	}
}

func (handler *AlertsHandler) GetAll(ctx *gin.Context) common.ApiError {
	response, err := handler.alertService.GetAll(ctx)
	if err != nil {
		return common.NewInternalServerApiError("Error getting alerts", err)
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}

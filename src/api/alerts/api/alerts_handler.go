package api

import (
	"challenge/alerts/src/api/alerts/domain"
	"challenge/alerts/src/api/application/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AlertsHandler struct {
	validate *validator.Validate
}

func NewAlertsHandler(validate *validator.Validate) AlertsHandler {
	return AlertsHandler{
		validate: validate,
	}
}

func (handler *AlertsHandler) GetAll(ctx *gin.Context) common.ApiError {
	response := []domain.AlertsDTO{
		{
			Type:        "xxxx",
			Description: "ddddd",
			CreatedAt:   "hh-mm-ss-dd-mm-yyyy",
			Country:     "Colombia",
		},
		{
			Type:        "yyyy",
			Description: "description example",
			CreatedAt:   "hh-mm-ss-dd-mm-yyyy",
			Country:     "Argentina",
		},
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}

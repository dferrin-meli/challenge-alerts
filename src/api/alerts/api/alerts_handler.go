package api

import (
	"challenge/alerts/src/api/alerts/domain"
	"challenge/alerts/src/api/application/common"
	"fmt"
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

func (handler *AlertsHandler) Create(ctx *gin.Context) common.ApiError {
	request := domain.AlertsDTO{}

	if err := ctx.BindQuery(&request); err != nil {
		return common.NewApiError("Error getting query data", err.Error(), http.StatusInternalServerError, common.CauseList{err})
	}

	if err := handler.validate.Struct(request); err != nil {
		errMessage := fmt.Sprintf("Invalid body for create alert: %s", err.Error())
		return common.NewBadRequestApiError(errMessage)
	}

	response, err := handler.alertService.Create(ctx, request)
	if err != nil {
		return common.NewInternalServerApiError("", err)
	}
	ctx.JSON(http.StatusCreated, response)
	return nil
}

func (handler *AlertsHandler) Search(ctx *gin.Context) common.ApiError {
	request := domain.AlertSearchDTO{}
	if err := ctx.BindQuery(&request); err != nil {
		return common.NewApiError("Error getting query data", err.Error(), http.StatusInternalServerError, common.CauseList{err})
	}

	if err := handler.validate.Struct(request); err != nil {
		errMessage := fmt.Sprintf("Invalid body for Search alert: %s", err.Error())
		return common.NewBadRequestApiError(errMessage)
	}
	response, err := handler.alertService.Search(ctx, request)
	if err != nil {
		return common.NewInternalServerApiError("Error getting alerts", err)
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}

func (handler *AlertsHandler) GetAlertByType(ctx *gin.Context) common.ApiError {
	request := domain.AlertSearchByTypeDTO{}
	if err := ctx.BindQuery(&request); err != nil {
		return common.NewApiError("Error getting query data", err.Error(), http.StatusInternalServerError, common.CauseList{err})
	}

	if err := handler.validate.Struct(request); err != nil {
		errMessage := fmt.Sprintf("Invalid body for Search alert by type: %s", err.Error())
		return common.NewBadRequestApiError(errMessage)
	}
	response, err := handler.alertService.GetAlertsByType(ctx, request)
	if err != nil {
		return common.NewInternalServerApiError("Error getting alerts", err)
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}

func (handler *AlertsHandler) GetMetrics(ctx *gin.Context) common.ApiError {
	response, err := handler.alertService.GetMetrics(ctx)
	if err != nil {
		return common.NewInternalServerApiError("Error getting alerts", err)
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}

package service

import (
	"challenge/alerts/src/api/alerts/domain"
	"challenge/alerts/src/api/application/common"
	"context"
)

type AlertsService struct {
	alertsRepository domain.AlertsRepository
}

func NewAlertsService(alertsRepository domain.AlertsRepository) domain.AlertsService {
	return &AlertsService{
		alertsRepository: alertsRepository,
	}
}

func (service *AlertsService) GetAll(ctx context.Context) ([]domain.AlertsDTO, error) {
	response, err := service.alertsRepository.GetAll(ctx)
	var alertsDto []domain.AlertsDTO
	for _, alert := range response {
		var alertDto domain.AlertsDTO
		alertDto.Country = alert.Country
		alertDto.CreatedAt = alert.CreatedAt
		alertDto.Description = alert.Description
		alertDto.Type = alert.Type

		alertsDto = append(alertsDto, alertDto)

	}
	if err != nil {
		// TODO log
		return nil, err // no return this error
	}
	return alertsDto, nil
}

func (service *AlertsService) Create(ctx context.Context, alert domain.AlertsDTO) (*domain.Alert, error) {
	common.NormalizeString(&alert.Country)
	common.NormalizeString(&alert.Type)
	response, err := service.alertsRepository.Create(ctx, domain.Alert(alert))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *AlertsService) Search(ctx context.Context, input domain.AlertSearchDTO) ([]domain.Alert, error) {
	response, err := service.alertsRepository.Search(ctx, domain.AlertSearch(input))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *AlertsService) GetAlertsByType(ctx context.Context, typeInput domain.AlertSearchByTypeDTO) ([]domain.Alert, error) {
	response, err := service.alertsRepository.GetAlertsByType(ctx, typeInput.Type)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *AlertsService) GetMetrics(ctx context.Context) ([]domain.Metrics, error) {
	response, err := service.alertsRepository.GetMetrics(ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}

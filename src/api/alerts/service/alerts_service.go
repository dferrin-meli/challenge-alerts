package service

import (
	"challenge/alerts/src/api/alerts/domain"
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
		return nil, err
	}
	return alertsDto, nil
}

package domain

import "context"

type AlertsDTO struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Country     string `json:"country"`
}

type Alerts struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Country     string `json:"country"`
}

type GetAllRequest struct {
	Description string `form:"description" binding:"max=50" validate:"omitempty,strregex=alpha_num"`
}

type AlertsService interface {
	GetAll(ctx context.Context) ([]AlertsDTO, error)
}

type AlertsRepository interface {
	GetAll(ctx context.Context) ([]Alerts, error)
}

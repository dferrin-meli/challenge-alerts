package domain

import "context"

type AlertsDTO struct {
	Type        string `form:"type" validate:"strregex=alpha_num"`
	Description string `form:"description" validate:"strregex=alpha_num"`
	CreatedAt   string `form:"created_at" binding:"required" validate:"strregex=datetime"`
	Country     string `form:"country" validate:"strregex=alpha_num"`
}

type AlertSearchDTO struct {
	Input string `form:"input" binding:"required" validate:"strregex=alpha_num"`
}

type AlertSearch struct {
	Input string
}

type Alert struct {
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
	Create(ctx context.Context, alert AlertsDTO) (*Alert, error)
	Search(ctx context.Context, input AlertSearchDTO) ([]Alert, error)
}

type AlertsRepository interface {
	GetAll(ctx context.Context) ([]Alert, error)
	Create(ctx context.Context, alert Alert) (*Alert, error)
	Search(ctx context.Context, input AlertSearch) ([]Alert, error)
}

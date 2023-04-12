package domain

type AlertsDTO struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Country     string `json:"country"`
}

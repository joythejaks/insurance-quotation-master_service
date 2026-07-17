package dto

type UpdateCountryRequest struct {
	Name     string `json:"name" binding:"max=100"`
	IsActive *bool  `json:"is_active"`
}

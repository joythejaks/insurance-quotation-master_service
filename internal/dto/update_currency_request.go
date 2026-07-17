package dto

type UpdateCurrencyRequest struct {
	Name     string `json:"name" binding:"max=100"`
	Symbol   string `json:"symbol" binding:"max=10"`
	IsActive *bool  `json:"is_active"`
}

package dto

type CreateCountryRequest struct {
	Code string `json:"code" binding:"required,max=10"`
	Name string `json:"name" binding:"required,max=100"`
}

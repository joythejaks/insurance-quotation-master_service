package dto

type CreateCurrencyRequest struct {
	Code   string `json:"code" binding:"required,max=10"`
	Name   string `json:"name" binding:"required,max=100"`
	Symbol string `json:"symbol" binding:"max=10"`
}

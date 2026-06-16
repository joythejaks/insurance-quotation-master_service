package dto

type CreateOccupationRequest struct {
	Code string `json:"code" binding:"required,max=50"`
	Name string `json:"name" binding:"required,max=255"`
}

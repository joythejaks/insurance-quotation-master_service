package dto

type UpdateOccupationRequest struct {
	Name     string `json:"name" binding:"max=255"`
	IsActive bool   `json:"is_active" binding:"required"`
}

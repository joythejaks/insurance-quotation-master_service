package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jordisetiawan/insurance-master-service/internal/dto"
	"github.com/jordisetiawan/insurance-master-service/internal/model"
	"github.com/jordisetiawan/insurance-master-service/internal/repository"
	"github.com/jordisetiawan/insurance-master-service/internal/utils"
	"gorm.io/gorm"
)

type OccupationHandler struct {
	repo repository.OccupationRepository
}

func NewOccupationHandler(repo repository.OccupationRepository) *OccupationHandler {
	return &OccupationHandler{repo: repo}
}

// @Summary Get all occupations
// @Tags Occupations
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param search query string false "Search by name or code"
// @Success 200 {object} utils.APIResponse{data=map[string]interface{}}
// @Router /occupations [get]
func (h *OccupationHandler) GetOccupations(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	occupations, total, err := h.repo.FindAll(page, limit, search)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch occupations", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Occupations fetched successfully", gin.H{
		"occupations": occupations,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

// @Summary Create a new occupation
// @Tags Occupations
// @Security BearerAuth
// @Accept json
// @Param occupation body dto.CreateOccupationRequest true "Occupation data"
// @Success 201 {object} utils.APIResponse{data=model.Occupation}
// @Router /occupations [post]
func (h *OccupationHandler) CreateOccupation(c *gin.Context) {
	var req dto.CreateOccupationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	occupation := model.Occupation{
		ID:       uuid.New(),
		Code:     req.Code,
		Name:     req.Name,
		IsActive: true,
	}

	if err := h.repo.Create(&occupation); err != nil {
		if err == gorm.ErrDuplicatedKey {
			utils.ErrorResponse(c, http.StatusConflict, "Occupation code already exists", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create occupation", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Occupation created successfully", occupation)
}

// @Summary Get a single occupation
// @Tags Occupations
// @Security BearerAuth
// @Param id path string true "Occupation ID"
// @Success 200 {object} utils.APIResponse{data=model.Occupation}
// @Router /occupations/{id} [get]
func (h *OccupationHandler) GetOccupation(c *gin.Context) {
	id := c.Param("id")
	occupation, err := h.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Occupation not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Occupation fetched successfully", occupation)
}

// @Summary Update an occupation
// @Tags Occupations
// @Security BearerAuth
// @Param id path string true "Occupation ID"
// @Accept json
// @Param occupation body dto.UpdateOccupationRequest true "Update data"
// @Success 200 {object} utils.APIResponse{data=model.Occupation}
// @Router /occupations/{id} [put]
func (h *OccupationHandler) UpdateOccupation(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateOccupationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	occupation, err := h.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Occupation not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	if req.Name != "" {
		occupation.Name = req.Name
	}
	occupation.IsActive = req.IsActive

	if err := h.repo.Update(occupation); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update occupation", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Occupation updated successfully", occupation)
}

// @Summary Delete an occupation
// @Tags Occupations
// @Security BearerAuth
// @Param id path string true "Occupation ID"
// @Success 200 {object} utils.APIResponse
// @Router /occupations/{id} [delete]
func (h *OccupationHandler) DeleteOccupation(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete occupation", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Occupation deleted successfully", nil)
}

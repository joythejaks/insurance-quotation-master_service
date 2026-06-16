package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordisetiawan/insurance-master-service/internal/dto"
	"github.com/jordisetiawan/insurance-master-service/internal/service"
	"github.com/jordisetiawan/insurance-master-service/internal/utils"
	"gorm.io/gorm"
)

type RelationshipHandler struct {
	service service.RelationshipService
}

func NewRelationshipHandler(
	service service.RelationshipService,
) *RelationshipHandler {
	return &RelationshipHandler{
		service: service,
	}
}

// @Summary Get all relationships
// @Tags Relationships
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param search query string false "Search by name or code"
// @Success 200 {object} utils.APIResponse{data=map[string]interface{}}
// @Router /relationships [get]
func (h *RelationshipHandler) GetRelationships(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	relationships, total, err := h.service.GetRelationships(
		page,
		limit,
		search,
	)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch relationships", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Relationships fetched successfully", gin.H{
		"relationships": relationships,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

// @Summary Create a new relationship
// @Tags Relationships
// @Security BearerAuth
// @Accept json
// @Param relationship body dto.CreateRelationshipRequest true "Relationship data"
// @Success 201 {object} utils.APIResponse{data=model.Relationship}
// @Router /relationships [post]
func (h *RelationshipHandler) CreateRelationship(c *gin.Context) {
	var req dto.CreateRelationshipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}
	relationship, err := h.service.CreateRelationship(req)

	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			utils.ErrorResponse(c, http.StatusConflict, "Relationship code already exists", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create relationship", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Relationship created successfully", relationship)
}

// @Summary Get a single relationship
// @Tags Relationships
// @Security BearerAuth
// @Param id path string true "Relationship ID"
// @Success 200 {object} utils.APIResponse{data=model.Relationship}
// @Router /relationships/{id} [get]
func (h *RelationshipHandler) GetRelationship(c *gin.Context) {
	id := c.Param("id")
	relationship, err := h.service.GetRelationship(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Relationship not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Relationship fetched successfully", relationship)
}

// @Summary Update a relationship
// @Tags Relationships
// @Security BearerAuth
// @Param id path string true "Relationship ID"
// @Accept json
// @Param relationship body dto.UpdateRelationshipRequest true "Update data"
// @Success 200 {object} utils.APIResponse{data=model.Relationship}
// @Router /relationships/{id} [put]
func (h *RelationshipHandler) UpdateRelationship(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateRelationshipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	relationship, err := h.service.UpdateRelationship(
		id,
		req,
	)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Relationship not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Relationship updated successfully", relationship)
}

// @Summary Delete a relationship
// @Tags Relationships
// @Security BearerAuth
// @Param id path string true "Relationship ID"
// @Success 200 {object} utils.APIResponse
// @Router /relationships/{id} [delete]
func (h *RelationshipHandler) DeleteRelationship(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteRelationship(id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete relationship", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Relationship deleted successfully", nil)
}

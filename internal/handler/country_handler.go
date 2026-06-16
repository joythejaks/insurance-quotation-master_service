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

type CountryHandler struct {
	service service.CountryService
}

func NewCountryHandler(
	service service.CountryService,
) *CountryHandler {
	return &CountryHandler{
		service: service,
	}
}

// @Summary Get all countries
// @Tags Countries
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param search query string false "Search by name or code"
// @Success 200 {object} utils.APIResponse{data=map[string]interface{}}
// @Router /countries [get]
func (h *CountryHandler) GetCountries(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	countries, total, err := h.service.GetCountries(
		page,
		limit,
		search,
	)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch countries", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Countries fetched successfully", gin.H{
		"countries": countries,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

// @Summary Create a new country
// @Tags Countries
// @Security BearerAuth
// @Accept json
// @Param country body dto.CreateCountryRequest true "Country data"
// @Success 201 {object} utils.APIResponse{data=model.Country}
// @Router /countries [post]
func (h *CountryHandler) CreateCountry(c *gin.Context) {
	var req dto.CreateCountryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	country, err := h.service.CreateCountry(req)
	if err != nil {
		// Handle unique constraint error for 'code'
		if err == gorm.ErrDuplicatedKey {
			utils.ErrorResponse(c, http.StatusConflict, "Country with this code already exists", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create country", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Country created successfully", country)
}

// @Summary Get a single country
// @Tags Countries
// @Security BearerAuth
// @Param id path string true "Country ID"
// @Success 200 {object} utils.APIResponse{data=model.Country}
// @Router /countries/{id} [get]
func (h *CountryHandler) GetCountry(c *gin.Context) {
	id := c.Param("id")
	country, err := h.service.GetCountry(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Country not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Country fetched successfully", country)
}

// @Summary Update a country
// @Tags Countries
// @Security BearerAuth
// @Param id path string true "Country ID"
// @Accept json
// @Param country body dto.UpdateCountryRequest true "Update data"
// @Success 200 {object} utils.APIResponse{data=model.Country}
// @Router /countries/{id} [put]
func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateCountryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	country, err := h.service.UpdateCountry(
		id,
		req,
	)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Country not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Country updated successfully", country)
}

// @Summary Delete a country
// @Tags Countries
// @Security BearerAuth
// @Param id path string true "Country ID"
// @Success 200 {object} utils.APIResponse
// @Router /countries/{id} [delete]
func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCountry(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Country not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Country deleted successfully", nil)
}

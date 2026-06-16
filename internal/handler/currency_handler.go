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

type CurrencyHandler struct {
	repo repository.CurrencyRepository
}

func NewCurrencyHandler(repo repository.CurrencyRepository) *CurrencyHandler {
	return &CurrencyHandler{repo: repo}
}

// @Summary Get all currencies
// @Tags Currencies
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param search query string false "Search by name or code"
// @Success 200 {object} utils.APIResponse{data=map[string]interface{}}
// @Router /currencies [get]
func (h *CurrencyHandler) GetCurrencies(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	currencies, total, err := h.repo.FindAll(page, limit, search)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch currencies", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Currencies fetched successfully", gin.H{
		"currencies": currencies,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

// @Summary Create a new currency
// @Tags Currencies
// @Security BearerAuth
// @Accept json
// @Param currency body dto.CreateCurrencyRequest true "Currency data"
// @Success 201 {object} utils.APIResponse{data=model.Currency}
// @Router /currencies [post]
func (h *CurrencyHandler) CreateCurrency(c *gin.Context) {
	var req dto.CreateCurrencyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	currency := model.Currency{
		ID:       uuid.New(),
		Code:     req.Code,
		Name:     req.Name,
		Symbol:   req.Symbol,
		IsActive: true,
	}

	if err := h.repo.Create(&currency); err != nil {
		if err == gorm.ErrDuplicatedKey {
			utils.ErrorResponse(c, http.StatusConflict, "Currency code already exists", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create currency", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Currency created successfully", currency)
}

// @Summary Get a single currency
// @Tags Currencies
// @Security BearerAuth
// @Param id path string true "Currency ID"
// @Success 200 {object} utils.APIResponse{data=model.Currency}
// @Router /currencies/{id} [get]
func (h *CurrencyHandler) GetCurrency(c *gin.Context) {
	id := c.Param("id")
	currency, err := h.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Currency not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Currency fetched successfully", currency)
}

// @Summary Update a currency
// @Tags Currencies
// @Security BearerAuth
// @Param id path string true "Currency ID"
// @Accept json
// @Param currency body dto.UpdateCurrencyRequest true "Update data"
// @Success 200 {object} utils.APIResponse{data=model.Currency}
// @Router /currencies/{id} [put]
func (h *CurrencyHandler) UpdateCurrency(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateCurrencyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	currency, err := h.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "Currency not found", nil)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	if req.Name != "" {
		currency.Name = req.Name
	}
	currency.Symbol = req.Symbol
	currency.IsActive = req.IsActive

	if err := h.repo.Update(currency); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update currency", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Currency updated successfully", currency)
}

// @Summary Delete a currency
// @Tags Currencies
// @Security BearerAuth
// @Param id path string true "Currency ID"
// @Success 200 {object} utils.APIResponse
// @Router /currencies/{id} [delete]
func (h *CurrencyHandler) DeleteCurrency(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete currency", err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "Currency deleted successfully", nil)
}

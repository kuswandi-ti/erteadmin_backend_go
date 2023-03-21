package handler

import (
	"erteadmin_backend/helper"
	"erteadmin_backend/lingkungan"
	"net/http"

	"github.com/gin-gonic/gin"
)

type lingkunganHandler struct {
	lingkunganService lingkungan.Service
}

func NewLingkunganHandler(service lingkungan.Service) *lingkunganHandler {
	return &lingkunganHandler{service}
}

func (h *lingkunganHandler) CreateLingkungan(c *gin.Context) {
	var input lingkungan.CreateLingkunganInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(err.Error(), http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newLingkungan, err := h.lingkunganService.SaveLingkungan(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := lingkungan.FormatLingkungan(newLingkungan)

	response := helper.APIResponse("Lingkungan sukses dibuat", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

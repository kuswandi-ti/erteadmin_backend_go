package handler

import (
	"erteadmin_backend/helper"
	"erteadmin_backend/wilayah"
	"net/http"

	"github.com/gin-gonic/gin"
)

type wilayahHandler struct {
	wilayahService wilayah.Service
}

func NewWilayahHandler(service wilayah.Service) *wilayahHandler {
	return &wilayahHandler{service}
}

func (h *wilayahHandler) GetAllProvinsi(c *gin.Context) {
	provinsi, err := h.wilayahService.GetAllProvinsi()
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data Provinsi", http.StatusOK, "success", wilayah.FormatProvinsiMany(provinsi))
	c.JSON(http.StatusOK, response)
}

// func (h *wilayahHandler) GetKabKotaByProvinsiID(c *gin.Context) {
// 	provinsiID, _ := strconv.Atoi(c.Param("provinsi_id"))

// 	kabkotas, err := h.wilayahService.GetKabKotaByProvinsiID(provinsiID)
// 	if err != nil {
// 		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Data Kabupaten & Kota", http.StatusOK, "success", wilayah.FormatKabKotaMany(kabkotas))
// 	c.JSON(http.StatusOK, response)
// }

func (h *wilayahHandler) GetKabKotaByProvinsiID(c *gin.Context) {
	var input wilayah.FindKabKotaInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Format error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	kabkotas, err := h.wilayahService.GetKabKotaByProvinsiID(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data Kabupaten & Kota", http.StatusOK, "success", wilayah.FormatKabKotaMany(kabkotas))
	c.JSON(http.StatusOK, response)
}

// func (h *wilayahHandler) GetKecamatanByKabKotaID(c *gin.Context) {
// 	kabkotaID, _ := strconv.Atoi(c.Param("kabkota_id"))

// 	kecamatans, err := h.wilayahService.GetKecamatanByKabKotaID(kabkotaID)
// 	if err != nil {
// 		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Data Kecamatan", http.StatusOK, "success", wilayah.FormatKecamatanMany(kecamatans))
// 	c.JSON(http.StatusOK, response)
// }

func (h *wilayahHandler) GetKecamatanByKabKotaID(c *gin.Context) {
	var input wilayah.FindKecamatanInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Format error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	kecamatans, err := h.wilayahService.GetKecamatanByKabKotaID(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data Kecamatan", http.StatusOK, "success", wilayah.FormatKecamatanMany(kecamatans))
	c.JSON(http.StatusOK, response)
}

func (h *wilayahHandler) GetKelurahanByKecamatanID(c *gin.Context) {
	var input wilayah.FindKelurahanInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Format error", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	kelurahans, err := h.wilayahService.GetKelurahanByKecamatanID(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data Kelurahan", http.StatusOK, "success", wilayah.FormatKelurahanMany(kelurahans))
	c.JSON(http.StatusOK, response)
}

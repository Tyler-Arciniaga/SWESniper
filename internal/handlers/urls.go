package handlers

import (
	"net/http"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	Service services.URLService
}

func (h *URLHandler) HandleAddURL(c *gin.Context) {
	var req models.AddURLRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}

	if e := h.Service.ValidateURLPost(&req); e != nil {
		c.JSON(http.StatusBadRequest, e.Error())
		return
	}

	err = h.Service.StoreURL(&req)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
	}

	c.Status(http.StatusCreated)
}

func (h *URLHandler) HandleGetURLs(c *gin.Context) {
	URLData, e := h.Service.GetAllURLs()

	if e != nil {
		c.JSON(http.StatusNotFound, e.Error())
	}

	c.JSON(http.StatusOK, URLData)
}

package handlers

import (
	"fmt"
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

	fmt.Print(req.URL, req.Description, req.CheckInterval)
	c.Status(http.StatusCreated)
}

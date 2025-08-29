package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func (h *HealthHandler) HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Healthy as can be :)",
	})
}

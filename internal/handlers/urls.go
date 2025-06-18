package handlers

import (
	"fmt"
	"net/http"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/gin-gonic/gin"
)

func HandleAddURL(c *gin.Context) {
	var req models.AddURLRequest
	err := c.BindJSON(&req)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	if valid := services.ValidateURLPost(&req); !valid {
		c.Status(http.StatusBadRequest)
		return
	}

	fmt.Print(req.URL, req.Description, req.CheckInterval)
	c.Status(http.StatusCreated)
}

package handlers

import (
	"net/http"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service services.AuthService
}

func (h *AuthHandler) HandleSignUp(c *gin.Context) {
	var user models.SignUpUser
	err := c.ShouldBindJSON(&user)
	//validate request body
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.Service.AddNewUser(&user)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

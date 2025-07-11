package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	//validate request body
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}

	//validate user authorization
	user, e := h.ExtractUserInfo(c)
	if e != nil {
		c.JSON(http.StatusUnauthorized, e)
	}

	log.Printf(" user: %v", user)

	//validate check interval
	if e := h.Service.ValidateURLPost(&req); e != nil {
		c.JSON(http.StatusBadRequest, e.Error())
		return
	}

	err = h.Service.StoreURL(&req, &user)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *URLHandler) ExtractUserInfo(c *gin.Context) (models.User, error) {
	authHeader := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	endpoint := fmt.Sprintf("https://%s.supabase.co/auth/v1/user", os.Getenv("SUPABASE_PROJECT_REF"))
	r, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	r.Header.Set("Authorization", "Bearer "+token)
	r.Header.Set("apikey", os.Getenv("SUPABASE_ANON_KEY"))

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil || resp.StatusCode != http.StatusOK {
		return models.User{}, err
	}

	var user models.User
	json.NewDecoder(resp.Body).Decode(&user)

	return user, nil

}

func (h *URLHandler) HandleGetURLs(c *gin.Context) {
	URLData, e := h.Service.GetAllURLs()

	if e != nil {
		c.JSON(http.StatusNotFound, e.Error())
		return
	}

	c.JSON(http.StatusOK, URLData)
}

func (h *URLHandler) HandleGetURLById(c *gin.Context) {
	urlID := c.Param("id")
	if urlID == "" {
		c.JSON(http.StatusBadRequest, "no urlID extracted from URL param")
		return
	}

	URLData, e := h.Service.GetURLById(urlID)

	if e != nil {
		c.JSON(http.StatusConflict, e.Error())
		return
	}

	c.JSON(http.StatusOK, URLData)
}

func (h *URLHandler) HandleDeleteURL(c *gin.Context) {
	urlID := c.Param("id")
	if urlID == "" {
		c.JSON(http.StatusBadRequest, "no urlID extracted from URL param")
		return
	}

	e := h.Service.DeleteURL(urlID)

	if e != nil {
		c.JSON(http.StatusConflict, e.Error())
		return
	}

	c.JSON(http.StatusOK, urlID)
}

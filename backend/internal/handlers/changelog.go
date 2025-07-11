package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/gin-gonic/gin"
)

type ChangeLogHandler struct {
	Service services.ChangeLogService
}

func (h *ChangeLogHandler) HandleGetAllChanges(c *gin.Context) {
	//extract user from request
	user, err := h.ExtractUserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}

	changeData, e := h.Service.GetAllChangeRecords(user)

	if e != nil {
		c.JSON(http.StatusInternalServerError, e.Error())
		return
	}

	c.JSON(http.StatusOK, changeData)
}

func (h *ChangeLogHandler) HandleGetURlChanges(c *gin.Context) {
	//extract user from request
	user, err := h.ExtractUserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}

	urlID := c.Param("id")

	if urlID == "" {
		c.JSON(http.StatusBadRequest, "no urlID extracted from URL param")
		return
	}

	changeData, e := h.Service.GetOneUrlChangeRecord(user, urlID)

	if e != nil {
		c.JSON(http.StatusConflict, e.Error())
		return
	}

	c.JSON(http.StatusOK, changeData)

}

func (h *ChangeLogHandler) ExtractUserInfo(c *gin.Context) (models.User, error) {
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

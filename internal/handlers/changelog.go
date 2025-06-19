package handlers

import (
	"net/http"

	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/gin-gonic/gin"
)

type ChangeLogHandler struct {
	Service services.ChangeLogService
}

func (h *ChangeLogHandler) HandleGetAllChanges(c *gin.Context) {
	changeData, e := h.Service.GetAllChangeRecords()

	if e != nil {
		c.JSON(http.StatusInternalServerError, e.Error())
	}

	c.JSON(http.StatusOK, changeData)
}

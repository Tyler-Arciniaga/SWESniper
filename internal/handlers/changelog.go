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
		return
	}

	c.JSON(http.StatusOK, changeData)
}

func (h *ChangeLogHandler) HandleGetURlChanges(c *gin.Context) {
	urlID := c.Param("id")

	if urlID == "" {
		c.JSON(http.StatusBadRequest, "no urlID extracted from URL param")
		return
	}

	changeData, e := h.Service.GetOneUrlChangeRecord(urlID)

	if e != nil {
		c.JSON(http.StatusConflict, e.Error())
		return
	}

	c.JSON(http.StatusOK, changeData)

}

package main

import (
	"github.com/Tyler-Arciniaga/SWESniper/internal/handlers"
	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/Tyler-Arciniaga/SWESniper/internal/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	urlStore := &storage.InMemStore{
		Data: make(map[string]models.URLRecord),
	}
	urlService := services.URLService{Store: urlStore}
	urlHandler := handlers.URLHandler{Service: urlService}

	router := gin.Default()
	router.POST("/urls", urlHandler.HandleAddURL)
	router.Run()
}

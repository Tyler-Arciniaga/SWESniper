package main

import (
	"github.com/Tyler-Arciniaga/SWESniper/internal/handlers"
	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/Tyler-Arciniaga/SWESniper/internal/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	//dev: instantiate url store, service and handler
	urlStore := &storage.InMemStore{
		Data: make(map[string]models.URLRecord),
	}
	urlService := services.URLService{Store: urlStore}
	urlHandler := handlers.URLHandler{Service: urlService}

	//create and start server with all proper handlers
	router := gin.Default()
	router.POST("/urls", urlHandler.HandleAddURL)
	//router.GET("/urls", urlHandler.HandleGetURLs)
	router.Run()
}

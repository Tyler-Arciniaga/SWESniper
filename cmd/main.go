package main

import (
	"github.com/Tyler-Arciniaga/SWESniper/internal/handlers"
	"github.com/Tyler-Arciniaga/SWESniper/internal/models"
	"github.com/Tyler-Arciniaga/SWESniper/internal/poller"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/Tyler-Arciniaga/SWESniper/internal/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	//dev: instantiate url store, service and handler
	inMemDB := &storage.InMemStore{
		URLTable:  make(map[string]models.URLRecord),
		ChangeLog: make(map[string][]models.ChangeRecord),
	}
	urlService := services.URLService{URLStore: inMemDB}
	urlHandler := handlers.URLHandler{Service: urlService}

	changeLogService := services.ChangeLogService{ChangeRepository: inMemDB}
	changeLogHandler := handlers.ChangeLogHandler{Service: changeLogService}

	scraperService := services.ScraperService{}

	poller := poller.Poller{UrlService: urlService, ChangeLogService: changeLogService, ScraperService: scraperService}

	go poller.StartPoller() //run poller in background (independent from req/res cycle)

	//create and start server with all proper handlers
	router := gin.Default()
	router.POST("/urls", urlHandler.HandleAddURL)
	router.GET("/urls", urlHandler.HandleGetURLs)
	router.GET("/changelog", changeLogHandler.HandleGetAllChanges)
	router.Run()
}

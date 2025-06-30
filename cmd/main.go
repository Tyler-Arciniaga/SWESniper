package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Tyler-Arciniaga/SWESniper/internal/handlers"
	"github.com/Tyler-Arciniaga/SWESniper/internal/notifier"
	"github.com/Tyler-Arciniaga/SWESniper/internal/poller"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/Tyler-Arciniaga/SWESniper/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	//create connection pool to postgres
	connstr := "postgresql://localhost/swesniper" //hardcoded databaseURL change later!!!
	dbpool, err := pgxpool.New(context.Background(), connstr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	/*
		//dev: instantiate url store, service and handler
		inMemDB := &storage.InMemStore{
			URLTable:  make(map[string]models.URLRecord),
			ChangeLog: make(map[string][]models.ChangeRecord),
		}
	*/

	db := &storage.Postgres{Pool: dbpool}
	urlService := services.URLService{URLStore: db}
	urlHandler := handlers.URLHandler{Service: urlService}

	changeLogService := services.ChangeLogService{ChangeRepository: db}
	changeLogHandler := handlers.ChangeLogHandler{Service: changeLogService}

	scraperService := services.ScraperService{}

	notifier := &notifier.BasicNotifier{}

	poller := poller.Poller{UrlService: urlService, ChangeLogService: changeLogService, ScraperService: scraperService, Notifier: notifier}

	go poller.StartPoller() //run poller in background (independent from req/res cycle)

	//create and start server with all proper handlers
	router := gin.Default()
	router.POST("/urls", urlHandler.HandleAddURL)
	router.GET("/urls", urlHandler.HandleGetURLs)
	router.GET("/urls/:id", urlHandler.HandleGetURLById)
	router.DELETE("/urls/:id", urlHandler.HandleDeleteURL)
	router.GET("/changelog", changeLogHandler.HandleGetAllChanges)
	router.GET("/changelog/:id", changeLogHandler.HandleGetURlChanges)
	router.Run()
}

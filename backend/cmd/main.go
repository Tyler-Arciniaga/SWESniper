package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Tyler-Arciniaga/SWESniper/internal/handlers"
	"github.com/Tyler-Arciniaga/SWESniper/internal/notifier"
	"github.com/Tyler-Arciniaga/SWESniper/internal/poller"
	"github.com/Tyler-Arciniaga/SWESniper/internal/services"
	"github.com/Tyler-Arciniaga/SWESniper/internal/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// CORS middleware
func main() {
	//load env variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	//connstr := "postgresql://localhost/swesniper" //hardcoded databaseURL change later!!!

	//extract connection string from environment variables
	connstr := os.Getenv("SUPABASE_API_KEY")
	if connstr == "" {
		fmt.Fprint(os.Stderr, "no connection string set")
		os.Exit(1)
	}

	//create a new config
	config, err := pgxpool.ParseConfig(connstr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connection string: %v\n", err)
		os.Exit(1)
	}

	//disable prepared statements to support Supabase transaction pooler
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	//create connection pool
	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
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

	db := &storage.Supabase{Pool: dbpool}
	urlService := services.URLService{URLStore: db}
	urlHandler := handlers.URLHandler{Service: urlService}

	changeLogService := services.ChangeLogService{ChangeRepository: db}
	changeLogHandler := handlers.ChangeLogHandler{Service: changeLogService}

	authService := services.AuthService{}
	authHandler := handlers.AuthHandler{Service: authService}

	scraperService := services.ScraperService{}

	notifier := &notifier.EmailNotifier{}

	poller := poller.Poller{UrlService: urlService, ChangeLogService: changeLogService, ScraperService: scraperService, Notifier: notifier}

	go poller.StartPoller() //run poller in background (independent from req/res cycle)

	//create and start server with all proper handlers
	router := gin.Default()

	//configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://swe-sniper.vercel.app"}, //PROD API
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/urls", urlHandler.HandleAddURL)
	router.GET("/urls", urlHandler.HandleGetURLs)
	router.GET("/urls/:id", urlHandler.HandleGetURLById)
	router.DELETE("/urls/:id", urlHandler.HandleDeleteURL)
	router.GET("/changelog", changeLogHandler.HandleGetAllChanges)
	router.GET("/changelog/:id", changeLogHandler.HandleGetURlChanges)
	router.POST("/signup", authHandler.HandleSignUp)
	router.Run()
}

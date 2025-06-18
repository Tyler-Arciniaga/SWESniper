package main

import (
	"github.com/Tyler-Arciniaga/SWESniper/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/urls", handlers.HandleAddURL)
	router.Run()
}

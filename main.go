package main

import (
	"hello/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", ".static")

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("server starting on port %s", port)
	r.Run("0.0.0.0:" + port)
}


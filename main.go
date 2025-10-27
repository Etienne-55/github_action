package main

import (
	"hello/routes"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("server starting on port %s", port)
	r.Run("0.0.0.0:" + port)
}

func test(t *testing.T) {
	if 2+2 != 4 {
		t.Error("Math broke!")
	}
}


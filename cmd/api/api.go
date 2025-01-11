package api

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/matthewjamesboyle/golang-interview-prep/internal/config"
	"github.com/matthewjamesboyle/golang-interview-prep/internal/db"
	routes "github.com/matthewjamesboyle/golang-interview-prep/internal/user"
)

var env = config.Envs

func StartServer() {
	db.StartDB()
	defer db.CloseDBConnection()
	server := gin.Default()

	// Define the API version group
	v1 := server.Group("/api/v1")
	{
		v1.GET("/test", routes.Test)
		v1.POST("/user", routes.CreateUser)
		v1.GET("/user/:id", routes.GetUserByID)
		v1.GET("/users", routes.GetUsers)
	}

	// Use Gin's built-in Run method to start the server on the desired port

	log.Printf("Starting HTTP server on http://localhost%s/api/v1/test\n", env.PORT)

	// Start the server using Gin's built-in method
	if err := server.Run(env.PORT); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

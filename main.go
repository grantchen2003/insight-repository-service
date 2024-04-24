package main

import (
	"fmt"
	"insight-repository-service/config"
	"insight-repository-service/database"
	"insight-repository-service/handlers"
	initializeRepository "insight-repository-service/handlers/initialize_repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("ENV")
	log.Printf("ENV=%s", env)
	envIsProduction, err := config.LoadEnvVars(env)

	if err != nil {
		log.Fatal(err)
		return
	}

	db := database.GetInstance()
	db.Connect()
	defer db.Close()

	if envIsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.POST("/initialize_repository", initializeRepository.InitializeRepository)
	router.GET("/query_repository", handlers.QueryRepository)
	router.PUT("/reinitialize_repository", handlers.ReinitializeRepository)
	router.DELETE("/uninitialize_repository", handlers.UninitializeRepository)
	router.POST("/validate_repository_id", handlers.ValidateRepositoryId)

	address := fmt.Sprintf("%s:%s", os.Getenv("DOMAIN"), os.Getenv("PORT"))
	log.Printf("Server running on %s", address)
	log.Fatal(router.Run(address))
}

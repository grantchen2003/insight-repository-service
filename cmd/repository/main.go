package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/config"
	"github.com/grantchen2003/insight/repository/internal/handlers"
	initializeRepository "github.com/grantchen2003/insight/repository/internal/handlers/initializerepository"

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

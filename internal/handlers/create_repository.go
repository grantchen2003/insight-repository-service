package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grantchen2003/insight/repository/internal/database"
)

func CreateRepository(c *gin.Context) {
	db := database.GetSingletonInstance()

	repositoryId, err := db.CreateRepository()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, map[string]string{
		"repository_id": repositoryId,
	})
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grantchen2003/insight/repository/internal/database"
)

type ValidateRepositoryIdRequest struct {
	RepositoryId string `json:"repository_id"`
}

func ValidateRepositoryId(c *gin.Context) {
	var request ValidateRepositoryIdRequest

	if err := c.BindJSON(&request); err != nil {
		panic(err)
	}

	db := database.GetSingletonInstance()

	repository, err := db.GetRepositoryById(request.RepositoryId)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, map[string]bool{
		"repository_id_is_valid": repository != nil,
	})
}

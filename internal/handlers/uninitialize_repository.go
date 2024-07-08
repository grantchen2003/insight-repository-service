package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grantchen2003/insight/repository/internal/database"
)

type UninitializeRepositoryRequest struct {
	RepositoryId string `json:"repository_id"`
}

func UninitializeRepository(c *gin.Context) {
	var request UninitializeRepositoryRequest

	if err := c.BindJSON(&request); err != nil {
		panic(err)
	}

	db := database.GetSingletonInstance()

	if err := db.DeleteRepository(request.RepositoryId); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, map[string]string{
		"repository_id": "id1",
	})
}

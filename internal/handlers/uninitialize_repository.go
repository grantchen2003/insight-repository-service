package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grantchen2003/insight/repository/internal/database"
	fileChunksService "github.com/grantchen2003/insight/repository/internal/services/filechunksservice"
	fileComponentService "github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice"
	vectorEmbedderService "github.com/grantchen2003/insight/repository/internal/services/vectorembedderservice"
)

type UninitializeRepositoryRequest struct {
	RepositoryId string `json:"repository_id"`
}

func UninitializeRepository(c *gin.Context) {
	var request UninitializeRepositoryRequest

	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	db := database.GetSingletonInstance()

	if err := db.DeleteRepository(request.RepositoryId); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	if err := fileChunksService.DeleteFileChunksByRepositoryId(request.RepositoryId); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	if err := fileComponentService.DeleteFileComponentsByRepositoryId(request.RepositoryId); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	if err := vectorEmbedderService.DeleteFileComponentVectorEmbeddingsByRepositoryId(request.RepositoryId); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	c.Status(http.StatusOK)
}

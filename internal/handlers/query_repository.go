package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	fileComponentService "github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice"
	vectorEmbedderService "github.com/grantchen2003/insight/repository/internal/services/vectorembedderservice"
)

type FileMatch struct {
	Path      string `json:"path"`
	StartLine int    `json:"start_line"`
	EndLine   int    `json:"end_line"`
	Content   string `json:"content"`
}

type QueryRepositoryRequest struct {
	RepositoryId string `json:"repository_id"`
	Query        string `json:"query_string"`
	Limit        int    `json:"limit"`
}

func QueryRepository(c *gin.Context) {
	var request QueryRepositoryRequest

	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	fileComponentIds, err := vectorEmbedderService.GetSimilarFileComponentIds(
		request.RepositoryId, request.Query, request.Limit,
	)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	fileComponents, err := fileComponentService.GetFileComponents(fileComponentIds)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	fileMatches := getFileMatches(fileComponents)

	c.JSON(http.StatusOK, fileMatches)
}

func getFileMatches(fileComponents []fileComponentService.FileComponent) []FileMatch {
	var fileMatches []FileMatch

	for _, fileComponent := range fileComponents {
		fileMatches = append(fileMatches, FileMatch{
			Path:      fileComponent.FilePath,
			StartLine: fileComponent.StartLine,
			EndLine:   fileComponent.EndLine,
			Content:   fileComponent.Content,
		})
	}

	return fileMatches
}

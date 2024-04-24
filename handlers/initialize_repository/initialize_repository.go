package initializeRepository

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func unpackRequest(c *gin.Context) (RepoInitBatch, error) {
	var batch RepoInitBatch
	if err := c.BindJSON(&batch); err != nil {
		return batch, err
	}

	sessionId, err := c.Cookie("session_id")
	if err != nil {
		return batch, err
	}

	batch.SessionId = sessionId
	return batch, err
}

func semanticallySummarizeFiles(repoId string, filePaths []string) {
}

func vectorEmbedFiles(repoId string, filePaths []string) {
}

func InitializeRepository(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"repository_id": "123",
	})

	batch, err := unpackRequest(c)
	if err != nil {
		return
	}

	_, err = batch.SaveRaw()
	if err != nil {
		return
	}

	// SyntaxParseFiles(batch.SessionId, savedFilePaths)
	// if err != nil {
	// 	return
	// }

	// semanticallySummarizeFiles(batch.SessionId, savedFilePaths)
	// if err != nil {
	// 	return
	// }

	// vectorEmbedFiles(batch.SessionId, savedFilePaths)
	// if err != nil {
	// 	return
	// }
}

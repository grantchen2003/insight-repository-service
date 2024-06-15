package initializeRepository

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	filecomponentservice "github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice"
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

func InitializeRepository(c *gin.Context) {
	// unpack the request
	batch, err := unpackRequest(c)
	if err != nil {
		// handle case where request couldn't be unpacked
		return
	}

	filePathsToProcess, err := batch.SaveFileChunks()
	if err != nil {
		// handle case where grpc report to lock doesn't work
		return
	}

	fileComponents, err := filecomponentservice.GetFilesComponents(batch.SessionId, filePathsToProcess)
	if err != nil {
		return
	}

	for _, fileComponent := range fileComponents {
		log.Println(fileComponent)
	}

	// semantically summarize each file in filePathsToProcess
	// vector embed each file in filePathsToProcess

	c.JSON(http.StatusOK, map[string]string{
		"repository_id": "123",
	})
}

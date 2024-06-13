package initializeRepository

import (
	fileSegmentService "insight-repository-service/services/filesegmentservice"
	"log"
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

func InitializeRepository(c *gin.Context) {
	// unpack the request
	batch, err := unpackRequest(c)
	if err != nil {
		// handle case where request couldn't be unpacked
		return
	}

	err = batch.SaveFileChunksAsBase64()
	if err != nil {
		// handle case where one/many/all chunks didn't save
		return
	}

	filePathsToProcess, err := batch.SyncFileChunks()
	if err != nil {
		// handle case where grpc report to lock doesn't work
		return
	}

	fileSegments, err := fileSegmentService.GetFileSegments(batch.SessionId, filePathsToProcess)
	if err != nil {
		return
	}

	for _, fileSegment := range fileSegments {
		log.Println(fileSegment)
	}

	// semantically summarize each file in filePathsToProcess
	// vector embed each file in filePathsToProcess

	c.JSON(http.StatusOK, map[string]string{
		"repository_id": "123",
	})
}

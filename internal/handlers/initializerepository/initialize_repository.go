package initializerepository

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	fileChunksService "github.com/grantchen2003/insight/repository/internal/services/filechunksservice"
	fileComponentService "github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice"
	vectorEmbedderService "github.com/grantchen2003/insight/repository/internal/services/vectorembedderservice"
)

type RepoInitBatch struct {
	SessionId string
	Files     map[string]struct {
		Content        string `json:"content"`
		ChunkIndex     int    `json:"chunk_index"`
		NumTotalChunks int    `json:"num_total_chunks"`
	} `json:"files"`
	BatchIndex      int `json:"batch_index"`
	NumTotalBatches int `json:"num_total_batches"`
}

func InitializeRepository(c *gin.Context) {
	batch, err := unpackRequest(c)
	if err != nil {
		return
	}

	fileChunks, err := batch.toFileChunks()
	if err != nil {
		return
	}

	fileChunkSaveStatuses, err := fileChunksService.CreateFileChunks(batch.SessionId, fileChunks)
	if err != nil {
		return
	}

	filePathsToProcess := getFilePathsToProcess(fileChunkSaveStatuses)

	fileComponents, err := fileComponentService.CreateFileComponents(batch.SessionId, filePathsToProcess)
	if err != nil {
		return
	}

	fileComponentIds := getFileComponentIds(fileComponents)

	if _, err := vectorEmbedderService.CreateFileComponentVectorEmbeddings(fileComponentIds); err != nil {
		log.Println(13)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"repository_id": "123",
	})
}

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

func getFilePathsToProcess(fileChunkSaveStatuses []fileChunksService.FileChunkSaveStatus) []string {
	var filePathsToProcess []string

	for _, fileChunkSaveStatus := range fileChunkSaveStatuses {
		if fileChunkSaveStatus.IsLastSavedChunk {
			filePathsToProcess = append(filePathsToProcess, fileChunkSaveStatus.FilePath)
		}
	}

	return filePathsToProcess
}

func (batch RepoInitBatch) toFileChunks() ([]fileChunksService.FileChunk, error) {
	var fileChunks []fileChunksService.FileChunk

	for filePath, fileData := range batch.Files {
		fileContent, err := base64.StdEncoding.DecodeString(fileData.Content)
		if err != nil {
			return nil, err
		}

		fileChunks = append(fileChunks, fileChunksService.FileChunk{
			UserId:         batch.SessionId,
			FilePath:       filePath,
			Content:        fileContent,
			ChunkIndex:     fileData.ChunkIndex,
			NumTotalChunks: fileData.NumTotalChunks,
		})
	}

	return fileChunks, nil
}

func getFileComponentIds(fileComponents []fileComponentService.FileComponent) []int32 {
	var fileComponentIds []int32

	for _, fileComponent := range fileComponents {
		fileComponentIds = append(fileComponentIds, int32(fileComponent.Id))
	}

	return fileComponentIds
}

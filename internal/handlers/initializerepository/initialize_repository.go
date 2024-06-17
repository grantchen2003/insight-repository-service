package initializerepository

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	fileChunksService "github.com/grantchen2003/insight/repository/internal/services/filechunksservice"
	fileComponentService "github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice"
	summarizerService "github.com/grantchen2003/insight/repository/internal/services/summarizerservice"
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

	fileChunkSaveStatuses, err := fileChunksService.SaveFileChunks(batch.SessionId, fileChunks)
	if err != nil {
		return
	}

	filePathsToProcess := getFilePathsToProcess(fileChunkSaveStatuses)

	fileComponents, err := fileComponentService.BatchExtractFileComponents(batch.SessionId, filePathsToProcess)
	if err != nil {
		return
	}

	fileComponentIds, err := fileComponentService.SaveFileComponents(fileComponents)
	if err != nil {
		return
	}

	fileComponentSummaries, err := summarizerService.CreateFileComponentSummaries(fileComponentIds)
	if err != nil {
		return
	}

	for _, fileComponentSummary := range fileComponentSummaries {
		fmt.Println(fileComponentSummary)
	}

	// fileComponentVectorEmbeddings, err := vectorEmbedderService.Embed(fileComponentSummaries)
	// fileComponentVectorEmbeddingIds, err := vectorEmbedderService.BatchSaveEmbeddings(fileComponentVectorEmbeddings)

	// for _, fileComponent := range fileComponents {
	// 	FileComponent{

	// 	}
	// }

	// semantically summarize each file in filePathsToProcess
	// vector embed each file in filePathsToProcess

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

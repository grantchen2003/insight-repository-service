package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	fileChunksService "github.com/grantchen2003/insight/repository/internal/services/filechunksservice"
	fileComponentService "github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice"
	vectorEmbedderService "github.com/grantchen2003/insight/repository/internal/services/vectorembedderservice"
)

type InitializeRepositoryRequest struct {
	RepositoryId string
	Files        map[string]struct {
		Content        string `json:"content"`
		ChunkIndex     int    `json:"chunk_index"`
		NumTotalChunks int    `json:"num_total_chunks"`
	} `json:"files"`
	BatchIndex      int `json:"batch_index"`
	NumTotalBatches int `json:"num_total_batches"`
}

func InitializeRepository(c *gin.Context) {
	request, err := unpackInitializeRepositoryRequest(c)
	if err != nil {
		return
	}

	fileChunks, err := getFileChunks(request)
	if err != nil {
		return
	}

	fileChunkSaveStatuses, err := fileChunksService.CreateFileChunks(request.RepositoryId, fileChunks)
	if err != nil {
		return
	}

	filePathsToProcess := getFilePathsToProcess(fileChunkSaveStatuses)

	fileComponents, err := fileComponentService.CreateFileComponents(request.RepositoryId, filePathsToProcess)
	if err != nil {
		return
	}

	fileComponentIds := getFileComponentIds(fileComponents)

	if _, err := vectorEmbedderService.CreateFileComponentVectorEmbeddings(fileComponentIds); err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"repository_id": "123",
	})
}

func unpackInitializeRepositoryRequest(c *gin.Context) (InitializeRepositoryRequest, error) {
	var request InitializeRepositoryRequest
	if err := c.BindJSON(&request); err != nil {
		return request, err
	}

	repositoryId, err := c.Cookie("repository_id")
	if err != nil {
		return request, err
	}

	request.RepositoryId = repositoryId
	return request, err
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

func getFileChunks(request InitializeRepositoryRequest) ([]fileChunksService.FileChunk, error) {
	var fileChunks []fileChunksService.FileChunk

	for filePath, fileData := range request.Files {
		fileContent, err := base64.StdEncoding.DecodeString(fileData.Content)
		if err != nil {
			return nil, err
		}

		fileChunks = append(fileChunks, fileChunksService.FileChunk{
			RepositoryId:   request.RepositoryId,
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

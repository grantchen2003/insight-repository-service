package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	fileChunksService "github.com/grantchen2003/insight/repository/internal/services/filechunksservice"
	fileComponentService "github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice"
	vectorEmbedderService "github.com/grantchen2003/insight/repository/internal/services/vectorembedderservice"
)

type ReinitializeRepositoryRequest struct {
	RepositoryId string `json:"repository_id"`
	Files        map[string]struct {
		Content        string `json:"content"`
		ChunkIndex     int    `json:"chunk_index"`
		NumTotalChunks int    `json:"num_total_chunks"`
	}
	Changes         map[string]string `json:"changes"`
	BatchIndex      int               `json:"batch_index"`
	NumTotalBatches int               `json:"num_total_batches"`
}

func ReinitializeRepository(c *gin.Context) {
	var request ReinitializeRepositoryRequest

	if err := c.BindJSON(&request); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	addedFilePaths, updatedFilePaths, deletedFilePaths := categorizeFileChanges(request.Changes)

	addedFileChunks, err := getFileChunksFromFilePaths(request, addedFilePaths)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	if err := addFiles(request.RepositoryId, addedFileChunks); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	updatedFileChunks, err := getFileChunksFromFilePaths(request, updatedFilePaths)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	if err := updateFiles(request.RepositoryId, updatedFileChunks); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	if err := deleteFiles(request.RepositoryId, deletedFilePaths); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	c.Status(http.StatusOK)
}

func categorizeFileChanges(changes map[string]string) ([]string, []string, []string) {
	var addedFilePaths []string
	var updatedFilePaths []string
	var deletedFilePaths []string

	for filePath, change := range changes {
		switch change {
		case "add":
			addedFilePaths = append(addedFilePaths, filePath)

		case "update":
			updatedFilePaths = append(updatedFilePaths, filePath)

		case "delete":
			deletedFilePaths = append(deletedFilePaths, filePath)
		}
	}

	return addedFilePaths, updatedFilePaths, deletedFilePaths
}

func deleteFiles(repositoryId string, filePaths []string) error {
	if len(filePaths) == 0 {
		return nil
	}

	err := fileChunksService.DeleteFileChunksByRepositoryIdAndFilePaths(repositoryId, filePaths)
	if err != nil {
		return err
	}

	fileComponentIds, err := fileComponentService.DeleteFileComponentsByRepositoryIdAndFilePaths(repositoryId, filePaths)
	if err != nil {
		return err
	}

	err = vectorEmbedderService.DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds(repositoryId, fileComponentIds)
	if err != nil {
		return err
	}

	return nil
}

func updateFiles(repositoryId string, fileChunks []fileChunksService.FileChunk) error {
	var filePaths []string
	for _, fileChunk := range fileChunks {
		filePaths = append(filePaths, fileChunk.FilePath)
	}

	if err := deleteFiles(repositoryId, filePaths); err != nil {
		return err
	}

	if err := addFiles(repositoryId, fileChunks); err != nil {
		return err
	}

	return nil
}

func getFileChunksFromFilePaths(request ReinitializeRepositoryRequest, filePaths []string) ([]fileChunksService.FileChunk, error) {
	var fileChunks []fileChunksService.FileChunk
	for _, filePath := range filePaths {
		fileChunk, err := getFileChunkFromFilePath(request, filePath)
		if err != nil {
			return nil, err
		}
		fileChunks = append(fileChunks, fileChunk)
	}
	return fileChunks, nil
}

func getFileChunkFromFilePath(request ReinitializeRepositoryRequest, filePath string) (fileChunksService.FileChunk, error) {
	fileContent, err := base64.StdEncoding.DecodeString(request.Files[filePath].Content)
	if err != nil {
		return fileChunksService.FileChunk{}, err
	}

	return fileChunksService.FileChunk{
		RepositoryId:   request.RepositoryId,
		FilePath:       filePath,
		Content:        fileContent,
		ChunkIndex:     request.Files[filePath].ChunkIndex,
		NumTotalChunks: request.Files[filePath].NumTotalChunks,
	}, nil
}

package initializeRepository

import (
	"log"

	fileChunksService "github.com/grantchen2003/insight/repository/internal/services/filechunksservice"
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

func (batch RepoInitBatch) SaveFileChunks() ([]string, error) {

	var fileChunks []fileChunksService.FileChunk

	for filePath, fileData := range batch.Files {
		fileChunks = append(fileChunks, fileChunksService.FileChunk{
			UserId:         batch.SessionId,
			FilePath:       filePath,
			Content:        fileData.Content,
			ChunkIndex:     fileData.ChunkIndex,
			NumTotalChunks: fileData.NumTotalChunks,
		})
	}

	fileChunkStatuses, err := fileChunksService.SaveFileChunks(batch.SessionId, fileChunks)

	if err != nil {
		log.Fatal(err)
	}

	var filePathsToProcess []string

	for _, fileChunkStatus := range fileChunkStatuses {
		log.Println(fileChunkStatus.FilePath, fileChunkStatus.IsLastSavedChunk)
		if fileChunkStatus.IsLastSavedChunk {
			filePathsToProcess = append(filePathsToProcess, fileChunkStatus.FilePath)
		}
	}

	return filePathsToProcess, nil
}

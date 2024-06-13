package initializeRepository

import (
	"insight-repository-service/database"
	repositorySyncService "insight-repository-service/services/repositorysyncservice"
	"log"
	"os"
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

type FileChunk struct {
	UserId            string
	FilePath          string
	Content           string
	ContentChunkIndex int
	NumTotalChunks    int
}

func (batch RepoInitBatch) SaveFileChunksAsBase64() error {
	var fileChunks []interface{}

	for filePath, fileData := range batch.Files {
		fileChunk := FileChunk{
			UserId:   batch.SessionId,
			FilePath: filePath,
			// content is stored as base64 string which can be decoded to give us a utf8 string
			Content:           fileData.Content,
			ContentChunkIndex: fileData.ChunkIndex,
			NumTotalChunks:    fileData.NumTotalChunks,
		}
		fileChunks = append(fileChunks, fileChunk)
	}

	db := database.GetInstance()

	return db.BatchSave(os.Getenv("FILE_CHUNKS_DB"), "file_chunks", fileChunks)
}

func (batch RepoInitBatch) SyncFileChunks() ([]string, error) {

	var fileChunks []repositorySyncService.FileChunk

	for filePath, fileData := range batch.Files {
		fileChunks = append(fileChunks, repositorySyncService.FileChunk{
			FilePath:       filePath,
			ChunkIndex:     fileData.ChunkIndex,
			NumTotalChunks: fileData.NumTotalChunks,
		})
	}

	fileChunkStatuses, err := repositorySyncService.SyncFileChunks(batch.SessionId, fileChunks)

	if err != nil {
		log.Fatal(err)
	}

	var filePathsToProcess []string

	for _, fileChunkStatus := range fileChunkStatuses {
		log.Println(fileChunkStatus.FilePath, fileChunkStatus.IsLastChunk)
		if fileChunkStatus.IsLastChunk {
			filePathsToProcess = append(filePathsToProcess, fileChunkStatus.FilePath)
		}
	}

	return filePathsToProcess, nil
}

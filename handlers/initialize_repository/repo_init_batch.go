package initializeRepository

import (
	"context"
	"insight-repository-service/database"
	"insight-repository-service/protobufs"
	"log"
	"os"

	"google.golang.org/grpc"
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

func (batch RepoInitBatch) ReportSavedFileChunks() ([]string, error) {
	var filePathsToProcess []string

	conn, err := grpc.Dial(os.Getenv("REPOSITORY_SYNC_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := protobufs.NewRepositorySyncServiceClient(conn)
	stream, err := client.ReportSavedFileChunk(context.Background())
	if err != nil {
		log.Fatalf("error creating stream: %v", err)
	}

	for filePath, fileData := range batch.Files {
		err := stream.Send(&protobufs.ReportSavedFileChunkRequest{
			UserId:         batch.SessionId,
			FilePath:       filePath,
			ChunkIndex:     int32(fileData.ChunkIndex),
			NumTotalChunks: int32(fileData.NumTotalChunks),
		})

		if err != nil {
			log.Fatalf("error sending request: %v", err)
		}

		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving response: %v", err)

		}

		if resp.IsLastChunk {
			filePathsToProcess = append(filePathsToProcess, filePath)
		}
	}

	if err = stream.CloseSend(); err != nil {
		log.Fatalf("error closing stream: %v", err)
	}

	return filePathsToProcess, nil
}

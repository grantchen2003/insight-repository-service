package initializeRepository

import (
	"context"
	"insight-repository-service/database"
	pb "insight-repository-service/protobufs"
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

func (batch RepoInitBatch) SaveRaw() ([]string, error) {
	var fileChunks []interface{}

	for filePath, fileData := range batch.Files {
		log.Printf(filePath)
		fileChunk := FileChunk{
			UserId:            batch.SessionId,
			FilePath:          filePath,
			Content:           fileData.Content,
			ContentChunkIndex: fileData.ChunkIndex,
			NumTotalChunks:    fileData.NumTotalChunks,
		}
		fileChunks = append(fileChunks, fileChunk)
	}

	db := database.GetInstance()

	err := db.BatchSave(os.Getenv("FILE_CHUNKS_DB"), "file_chunks", fileChunks)

	savedFilePaths := []string{}

	if err != nil {
		batch.ReportSaveRawFailure(err)
	} else {
		savedFilePaths, err = batch.ReportSaveRawSuccess()
	}

	return savedFilePaths, err
}

func (batch RepoInitBatch) ReportSaveRawFailure(err error) error {
	// grpc call to server
	return nil
}

// make this concurrent
func (batch RepoInitBatch) ReportSaveRawSuccess() ([]string, error) {
	fullySavedFilePaths := []string{}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRepositoryLockClient(conn)

	stream, err := client.HandleSaveRawSuccess(context.Background())
	if err != nil {
		return fullySavedFilePaths, err
	}

	for filePath, fileData := range batch.Files {
		err := stream.Send(&pb.ReportSaveRawSuccessRequest{
			UserId:            batch.SessionId,
			FilePath:          filePath,
			ContentChunkIndex: int32(fileData.ChunkIndex),
			NumTotalChunks:    int32(fileData.NumTotalChunks),
		})

		if err != nil {
			return fullySavedFilePaths, err
		}

		res, err := stream.Recv()
		if err != nil {
			return fullySavedFilePaths, err
		}

		if res.AllChunksSaved {
			fullySavedFilePaths = append(fullySavedFilePaths, res.FilePath)
		}
	}
	return fullySavedFilePaths, nil
}

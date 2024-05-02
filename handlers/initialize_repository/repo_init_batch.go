package initializeRepository

import (
	"insight-repository-service/database"
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
			// content is stored as base64 random string,
			// this base64 string can be decoded to give us a utf8 string
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

	for filePath, fileData := range batch.Files {
		// FILL ME IN
		// send the below data to the grpc server
		// make concurrent (bidirectional streaming?)
		log.Println(batch.SessionId, filePath, fileData.ChunkIndex, fileData.NumTotalChunks)
		isLastReportedChunkForFile := true
		if isLastReportedChunkForFile {
			filePathsToProcess = append(filePathsToProcess, filePath)
		}
	}
	return filePathsToProcess, nil
}

// make this concurrent
// func (batch RepoInitBatch) ReportSaveRawSuccess() ([]string, error) {
// 	fullySavedFilePaths := []string{}

// 	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("failed to dial: %v", err)
// 	}
// 	defer conn.Close()
// 	client := pb.NewRepositoryLockClient(conn)

// 	stream, err := client.HandleSaveRawSuccess(context.Background())
// 	if err != nil {
// 		return fullySavedFilePaths, err
// 	}

// 	for filePath, fileData := range batch.Files {
// 		err := stream.Send(&pb.ReportSaveRawSuccessRequest{
// 			UserId:            batch.SessionId,
// 			FilePath:          filePath,
// 			ContentChunkIndex: int32(fileData.ChunkIndex),
// 			NumTotalChunks:    int32(fileData.NumTotalChunks),
// 		})

// 		if err != nil {
// 			return fullySavedFilePaths, err
// 		}

// 		res, err := stream.Recv()
// 		if err != nil {
// 			return fullySavedFilePaths, err
// 		}

// 		if res.AllChunksSaved {
// 			fullySavedFilePaths = append(fullySavedFilePaths, res.FilePath)
// 		}
// 	}
// 	return fullySavedFilePaths, nil
// }

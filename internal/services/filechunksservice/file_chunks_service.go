package repositorysyncservice

import (
	"context"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/filechunksservice/pb"

	"google.golang.org/grpc"
)

type FileChunk struct {
	UserId         string
	FilePath       string
	Content        string
	ChunkIndex     int
	NumTotalChunks int
}

type FileChunkStatus struct {
	FilePath         string
	IsLastSavedChunk bool
}

func SaveFileChunks(userId string, fileChunks []FileChunk) ([]FileChunkStatus, error) {
	conn, err := grpc.Dial(os.Getenv("REPOSITORY_SYNC_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	var pbFileChunks []*pb.FileChunk

	for _, fileChunk := range fileChunks {
		pbFileChunks = append(pbFileChunks, &pb.FileChunk{
			UserId:         userId,
			FilePath:       fileChunk.FilePath,
			Content:        fileChunk.Content,
			ChunkIndex:     int32(fileChunk.ChunkIndex),
			NumTotalChunks: int32(fileChunk.NumTotalChunks),
		})
	}

	client := pb.NewFileChunksServiceClient(conn)
	response, err := client.SaveFileChunks(context.Background(), &pb.SaveFileChunksRequest{
		FileChunks: pbFileChunks,
	})

	if err != nil {
		log.Fatalf("Failed to syn file chunks: %v", err)
	}

	var fileChunkStatuses []FileChunkStatus

	for _, pbFileChunkStatus := range response.FileChunkStatuses {
		fileChunkStatuses = append(fileChunkStatuses, FileChunkStatus{
			FilePath:         pbFileChunkStatus.FilePath,
			IsLastSavedChunk: pbFileChunkStatus.IsLastSavedChunk,
		})
	}

	return fileChunkStatuses, nil

}

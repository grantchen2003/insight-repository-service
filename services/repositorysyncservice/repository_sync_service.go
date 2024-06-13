package repositorysyncservice

import (
	"context"
	"insight-repository-service/services/repositorysyncservice/pb"
	"log"
	"os"

	"google.golang.org/grpc"
)

type FileChunk struct {
	FilePath       string
	ChunkIndex     int
	NumTotalChunks int
}

type FileChunkStatus struct {
	FilePath    string
	IsLastChunk bool
}

func SyncFileChunks(userId string, fileChunks []FileChunk) ([]FileChunkStatus, error) {
	conn, err := grpc.Dial(os.Getenv("REPOSITORY_SYNC_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	var pbFileChunks []*pb.FileChunk

	for _, fileChunk := range fileChunks {
		pbFileChunks = append(pbFileChunks, &pb.FileChunk{
			FilePath:       fileChunk.FilePath,
			ChunkIndex:     int32(fileChunk.ChunkIndex),
			NumTotalChunks: int32(fileChunk.NumTotalChunks),
		})
	}

	client := pb.NewRepositorySyncServiceClient(conn)
	pbFileChunkStatuses, err := client.SyncFileChunks(context.Background(), &pb.FileChunks{
		UserId:     userId,
		FileChunks: pbFileChunks,
	})

	if err != nil {
		log.Fatalf("Failed to syn file chunks: %v", err)
	}

	var fileChunkStatuses []FileChunkStatus

	for _, pbFileChunkStatus := range pbFileChunkStatuses.FileChunkStatus {
		fileChunkStatuses = append(fileChunkStatuses, FileChunkStatus{
			FilePath:    pbFileChunkStatus.FilePath,
			IsLastChunk: pbFileChunkStatus.IsLastChunk,
		})
	}

	return fileChunkStatuses, nil

}

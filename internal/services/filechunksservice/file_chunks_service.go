package repositorysyncservice

import (
	"context"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/filechunksservice/pb"

	"google.golang.org/grpc"
)

type FileChunk struct {
	RepositoryId   string
	FilePath       string
	Content        []byte
	ChunkIndex     int
	NumTotalChunks int
}

type FileChunkSaveStatus struct {
	FilePath         string
	IsLastSavedChunk bool
}

func CreateFileChunks(repositoryId string, fileChunks []FileChunk) ([]FileChunkSaveStatus, error) {
	conn, err := grpc.Dial(os.Getenv("FILE_CHUNKS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var pbFileChunks []*pb.FileChunkPayload

	for _, fileChunk := range fileChunks {
		pbFileChunks = append(pbFileChunks, &pb.FileChunkPayload{
			RepositoryId:   repositoryId,
			FilePath:       fileChunk.FilePath,
			Content:        fileChunk.Content,
			ChunkIndex:     int32(fileChunk.ChunkIndex),
			NumTotalChunks: int32(fileChunk.NumTotalChunks),
		})
	}

	client := pb.NewFileChunksServiceClient(conn)
	response, err := client.CreateFileChunks(context.Background(), &pb.CreateFileChunksRequest{
		FileChunkPayloads: pbFileChunks,
	})

	if err != nil {
		return nil, err
	}

	var fileChunkSaveStatuses []FileChunkSaveStatus

	for _, pbFileChunkStatus := range response.FileChunkStatuses {
		fileChunkSaveStatuses = append(fileChunkSaveStatuses, FileChunkSaveStatus{
			FilePath:         pbFileChunkStatus.FilePath,
			IsLastSavedChunk: pbFileChunkStatus.IsLastSavedChunk,
		})
	}

	return fileChunkSaveStatuses, nil
}

func DeleteFileChunksByRepositoryId(repositoryId string) error {
	conn, err := grpc.Dial(os.Getenv("FILE_CHUNKS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewFileChunksServiceClient(conn)
	_, err = client.DeleteFileChunksByRepositoryId(context.Background(), &pb.DeleteFileChunksByRepositoryIdRequest{
		RepositoryId: repositoryId,
	})

	return err
}

func DeleteFileChunksByRepositoryIdAndFilePaths(repositoryId string, filePaths []string) error {
	conn, err := grpc.Dial(os.Getenv("FILE_CHUNKS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewFileChunksServiceClient(conn)
	_, err = client.DeleteFileChunksByRepositoryIdAndFilePaths(context.Background(), &pb.DeleteFileChunksByRepositoryIdAndFilePathsRequest{
		RepositoryId: repositoryId,
		FilePaths:    filePaths,
	})

	return err
}

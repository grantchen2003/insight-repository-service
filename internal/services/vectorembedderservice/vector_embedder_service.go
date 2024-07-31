package vectorembedderservice

import (
	"context"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/vectorembedderservice/pb"
	"google.golang.org/grpc"
)

func CreateFileComponentVectorEmbeddings(fileComponentIds []int32) ([]string, error) {
	conn, err := grpc.Dial(os.Getenv("VECTOR_EMBEDDER_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to vector embedder service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewVectorEmbedderServiceClient(conn)
	request := &pb.CreateFileComponentVectorEmbeddingsRequest{
		FileComponentIds: fileComponentIds,
	}

	resp, err := client.CreateFileComponentVectorEmbeddings(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return resp.FileComponentVectorEmbeddingIds, nil
}

func GetSimilarFileComponentIds(repositoryId string, query string, limit int) ([]int32, error) {
	conn, err := grpc.Dial(os.Getenv("VECTOR_EMBEDDER_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to vector embedder service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewVectorEmbedderServiceClient(conn)
	request := &pb.GetSimilarFileComponentIdsRequest{
		RepositoryId: repositoryId,
		Query:        query,
		Limit:        int32(limit),
	}

	resp, err := client.GetSimilarFileComponentIds(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return resp.FileComponentIds, nil
}

func DeleteFileComponentVectorEmbeddingsByRepositoryId(repositoryId string) error {
	conn, err := grpc.Dial(os.Getenv("VECTOR_EMBEDDER_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to vector embedder service: %v", err)
		return err
	}
	defer conn.Close()

	client := pb.NewVectorEmbedderServiceClient(conn)
	request := &pb.DeleteFileComponentVectorEmbeddingsByRepositoryIdRequest{RepositoryId: repositoryId}

	_, err = client.DeleteFileComponentVectorEmbeddingsByRepositoryId(context.Background(), request)
	return err
}

func DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds(repositoryId string, fileComponentIds []int32) error {
	conn, err := grpc.Dial(os.Getenv("VECTOR_EMBEDDER_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to vector embedder service: %v", err)
		return err
	}
	defer conn.Close()

	client := pb.NewVectorEmbedderServiceClient(conn)
	request := &pb.DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIdsRequest{
		RepositoryId:     repositoryId,
		FileComponentIds: fileComponentIds,
	}

	_, err = client.DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds(context.Background(), request)
	return err
}

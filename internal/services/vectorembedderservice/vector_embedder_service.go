package vectorembedderservice

import (
	"context"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/vectorembedderservice/pb"
	"google.golang.org/grpc"
)

func CreateFileComponentVectorEmbeddings(fileComponentIds []int32) ([]int32, error) {
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

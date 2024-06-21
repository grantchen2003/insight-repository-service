package summarizerservice

import (
	"context"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/summarizerservice/pb"
	"google.golang.org/grpc"
)

type FileComponentSummary struct {
	Id              int
	FileComponentId int
	Summary         string
}

func CreateFileComponentSummaries(fileComponentIds []int32) ([]FileComponentSummary, error) {
	conn, err := grpc.Dial(os.Getenv("SUMMARIZER_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to summarizer service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewSummarizerServiceClient(conn)
	request := &pb.CreateFileComponentSummariesRequest{
		FileComponentIds: fileComponentIds,
	}

	resp, err := client.CreateFileComponentSummaries(context.Background(), request)
	if err != nil {
		return nil, err
	}

	var fileComponentSummaries []FileComponentSummary

	for _, fileComponentSummary := range resp.FileComponentSummaries {
		fileComponentSummaries = append(fileComponentSummaries, FileComponentSummary{
			Id:              int(fileComponentSummary.Id),
			FileComponentId: int(fileComponentSummary.FileComponentId),
			Summary:         fileComponentSummary.Summary,
		})
	}

	return fileComponentSummaries, nil
}

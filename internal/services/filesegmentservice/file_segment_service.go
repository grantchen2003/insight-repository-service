package filesegmentservice

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/filesegmentservice/pb"

	"google.golang.org/grpc"
)

type FileSegment struct {
	filePath  string
	startLine int
	endLine   int
}

func GetFileSegments(userId string, filePaths []string) ([]FileSegment, error) {
	var fileSegments []FileSegment

	conn, err := grpc.Dial(os.Getenv("FILE_SEGMENT_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return fileSegments, err
	}
	defer conn.Close()

	client := pb.NewFileSegmentServiceClient(conn)
	request := &pb.Files{UserId: userId, FilePaths: filePaths}

	stream, err := client.ExtractStructure(context.Background(), request)

	if err != nil {
		log.Fatalf("error calling file segment service: %v", err)
		return fileSegments, err
	}

	for {
		fileSegment, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error receiving response from file segment service: %v", err)
			return fileSegments, err
		}

		fileSegments = append(fileSegments, FileSegment{
			filePath:  fileSegment.FilePath,
			startLine: int(fileSegment.StartLine),
			endLine:   int(fileSegment.EndLine),
		})
	}

	return fileSegments, err
}

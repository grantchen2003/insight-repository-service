package filecomponentservice

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice/pb"
	"google.golang.org/grpc"
)

type FileComponent struct {
	filePath  string
	startLine int
	endLine   int
}

func GetFilesComponents(userId string, filePaths []string) ([]FileComponent, error) {
	var fileComponents []FileComponent

	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.Files{UserId: userId, FilePaths: filePaths}

	stream, err := client.ExtractFilesComponents(context.Background(), request)

	if err != nil {
		log.Fatalf("error calling file component service: %v", err)
		return nil, err
	}

	for {
		fileComponent, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error receiving response from file segment service: %v", err)
			return nil, err
		}

		fileComponents = append(fileComponents, FileComponent{
			filePath:  fileComponent.FilePath,
			startLine: int(fileComponent.StartLine),
			endLine:   int(fileComponent.EndLine),
		})
	}

	return fileComponents, nil
}

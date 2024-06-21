package filecomponentservice

import (
	"context"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice/pb"
	"google.golang.org/grpc"
)

type FileComponent struct {
	Id       int
	UserId   string
	FilePath string
	Content  string
}

type FileComponentIds []int32

func CreateFileComponents(userId string, filePaths []string) ([]FileComponent, error) {
	var fileComponents []FileComponent

	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to file components service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.UserFilePaths{UserId: userId, FilePaths: filePaths}

	resp, err := client.CreateFileComponents(context.Background(), request)

	if err != nil {
		return nil, err
	}

	for _, fileComponent := range resp.FileComponents {
		fileComponents = append(fileComponents, FileComponent{
			Id:       int(fileComponent.Id),
			UserId:   fileComponent.UserId,
			FilePath: fileComponent.FilePath,
			Content:  fileComponent.Content,
		})
	}

	return fileComponents, nil
}

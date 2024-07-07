package filecomponentservice

import (
	"context"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice/pb"
	"google.golang.org/grpc"
)

type FileComponent struct {
	Id           int
	RepositoryId string
	FilePath     string
	StartLine    int
	EndLine      int
	Content      string
}

type FileComponentIds []int32

func CreateFileComponents(repositoryId string, filePaths []string) ([]FileComponent, error) {
	var fileComponents []FileComponent

	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to file components service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.RepositoryFilePaths{RepositoryId: repositoryId, FilePaths: filePaths}

	resp, err := client.CreateFileComponents(context.Background(), request)

	if err != nil {
		return nil, err
	}

	for _, fileComponent := range resp.FileComponents {
		fileComponents = append(fileComponents, FileComponent{
			Id:           int(fileComponent.Id),
			RepositoryId: fileComponent.RepositoryId,
			FilePath:     fileComponent.FilePath,
			StartLine:    int(fileComponent.StartLine),
			EndLine:      int(fileComponent.EndLine),
			Content:      fileComponent.Content,
		})
	}

	return fileComponents, nil
}

func GetFileComponents(fileComponentIds []int32) ([]FileComponent, error) {
	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to file components service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.FileComponentIds{FileComponentIds: fileComponentIds}

	resp, err := client.GetFileComponents(context.Background(), request)

	if err != nil {
		return nil, err
	}

	var fileComponents []FileComponent

	for _, fileComponent := range resp.FileComponents {
		fileComponents = append(fileComponents, FileComponent{
			Id:           int(fileComponent.Id),
			RepositoryId: fileComponent.RepositoryId,
			FilePath:     fileComponent.FilePath,
			StartLine:    int(fileComponent.StartLine),
			EndLine:      int(fileComponent.EndLine),
			Content:      fileComponent.Content,
		})
	}

	return fileComponents, nil
}

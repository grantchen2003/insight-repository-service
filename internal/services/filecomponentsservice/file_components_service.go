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
		panic(err)
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.RepositoryFilePaths{RepositoryId: repositoryId, FilePaths: filePaths}

	stream, err := client.CreateFileComponents(context.Background(), request)
	if err != nil {
		panic(err)
	}

	for {
		fileComponent, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

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
		panic(err)
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.FileComponentIds{FileComponentIds: fileComponentIds}

	stream, err := client.GetFileComponents(context.Background(), request)
	if err != nil {
		panic(err)
	}

	var fileComponents []FileComponent

	for {
		fileComponent, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

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

func DeleteFileComponentsByRepositoryId(repositoryId string) error {
	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to file components service: %v", err)
		return err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.DeleteFileComponentsByRepositoryIdRequest{RepositoryId: repositoryId}

	_, err = client.DeleteFileComponentsByRepositoryId(context.Background(), request)

	return err
}

func DeleteFileComponentsByRepositoryIdAndFilePaths(repositoryId string, filePaths []string) ([]int32, error) {
	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to file components service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.DeleteFileComponentsByRepositoryIdAndFilePathsRequest{
		RepositoryId: repositoryId,
		FilePaths:    filePaths,
	}

	response, err := client.DeleteFileComponentsByRepositoryIdAndFilePaths(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return response.FileComponentIds, nil
}

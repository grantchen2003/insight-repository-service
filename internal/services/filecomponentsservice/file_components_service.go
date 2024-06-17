package filecomponentservice

import (
	"context"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/filecomponentsservice/pb"
	"google.golang.org/grpc"
)

type FileComponent struct {
	userId    string
	filePath  string
	startLine int
	endLine   int
}

type SavedFileComponent struct {
	id        int32
	userId    string
	filePath  string
	startLine int
	endLine   int
}

type SavedFileComponentIds []int32

func GetSavedFileComponents(savedFileComponentIds SavedFileComponentIds) ([]SavedFileComponent, error) {
	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to file components service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.SavedFileComponentIds{
		SavedFileComponentIds: savedFileComponentIds,
	}

	resp, err := client.GetSavedFileComponents(context.Background(), request)
	if err != nil {
		return nil, err
	}

	var savedFileComponents []SavedFileComponent
	for _, pbSavedFileComponent := range resp.SavedFileComponents {
		savedFileComponents = append(savedFileComponents, SavedFileComponent{
			id:        pbSavedFileComponent.Id,
			userId:    pbSavedFileComponent.UserId,
			filePath:  pbSavedFileComponent.FilePath,
			startLine: int(pbSavedFileComponent.StartLine),
			endLine:   int(pbSavedFileComponent.EndLine),
		})
	}

	return savedFileComponents, nil
}

func SaveFileComponents(fileComponents []FileComponent) (SavedFileComponentIds, error) {

	var pbFileComponents []*pb.FileComponent
	for _, fileComponent := range fileComponents {
		pbFileComponents = append(pbFileComponents, &pb.FileComponent{
			UserId:    fileComponent.userId,
			FilePath:  fileComponent.filePath,
			StartLine: int32(fileComponent.startLine),
			EndLine:   int32(fileComponent.endLine),
		})
	}

	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to file components service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.FileComponents{
		FileComponents: pbFileComponents,
	}

	resp, err := client.SaveFileComponents(context.Background(), request)

	if err != nil {
		return nil, err
	}

	return resp.SavedFileComponentIds, nil

}

func BatchExtractFileComponents(userId string, filePaths []string) ([]FileComponent, error) {
	var fileComponents []FileComponent

	conn, err := grpc.Dial(os.Getenv("FILE_COMPONENTS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to file components service: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewFileComponentsServiceClient(conn)
	request := &pb.BatchExtractFileComponentsRequest{UserId: userId, FilePaths: filePaths}

	resp, err := client.BatchExtractFileComponents(context.Background(), request)

	if err != nil {
		return nil, err
	}

	for _, fileComponent := range resp.FileComponents {
		fileComponents = append(fileComponents, FileComponent{
			userId:    fileComponent.UserId,
			filePath:  fileComponent.FilePath,
			startLine: int(fileComponent.StartLine),
			endLine:   int(fileComponent.EndLine),
		})
	}

	return fileComponents, nil
}

package initializeRepository

import (
	"context"
	"insight-repository-service/protobufs"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func unpackRequest(c *gin.Context) (RepoInitBatch, error) {
	var batch RepoInitBatch
	if err := c.BindJSON(&batch); err != nil {
		return batch, err
	}

	sessionId, err := c.Cookie("session_id")
	if err != nil {
		return batch, err
	}

	batch.SessionId = sessionId
	return batch, err
}

func getFileStructure(userId string, filePaths []string) error {
	conn, err := grpc.Dial(os.Getenv("FILE_STRUCTURE_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := protobufs.NewFileStructureServiceClient(conn)
	request := &protobufs.Files{UserId: userId, FilePaths: filePaths}

	stream, err := client.ExtractStructure(context.Background(), request)

	if err != nil {
		log.Fatalf("error calling file structure service: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error receiving response from file structure service: %v", err)
		}

		log.Printf(response.FilePath)
	}

	return err

}

func InitializeRepository(c *gin.Context) {
	// unpack the request
	batch, err := unpackRequest(c)
	if err != nil {
		// handle case where request couldn't be unpacked
		return
	}

	err = batch.SaveFileChunksAsBase64()
	if err != nil {
		// handle case where one/many/all chunks didn't save
		return
	}

	filePathsToProcess, err := batch.ReportSavedFileChunks()
	if err != nil {
		// handle case where grpc report to lock doesn't work
		return
	}

	getFileStructure(batch.SessionId, filePathsToProcess)

	// syntax parse each file in filePathsToProcess
	// semantically summarize each file in filePathsToProcess
	// vector embed each file in filePathsToProcess

	c.JSON(http.StatusOK, map[string]string{
		"repository_id": "123",
	})
}

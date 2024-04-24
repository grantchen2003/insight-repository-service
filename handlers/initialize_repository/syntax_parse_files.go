package initializeRepository

import (
	"context"
	pb "insight-repository-service/protobufs"
	"io"
	"log"

	"google.golang.org/grpc"
)

func SyntaxParseFiles(userId string, filePaths []string) error {
	conn, err := grpc.Dial("localhost:50057", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	client := pb.NewSourceCodeParserClient(conn)

	stream, err := client.SyntaxParse(context.Background())
	if err != nil {
		return err
	}

	log.Println(filePaths)
	for _, filePath := range filePaths {
		log.Println(filePath)
		err := stream.Send(&pb.SyntaxParseRequest{
			UserId:   userId,
			FilePath: filePath,
		})

		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	// Close the sending side of the stream
	if err := stream.CloseSend(); err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("here")
			log.Fatal(err)
			return err
		}

		log.Println(res.Success)
	}

	return nil
}

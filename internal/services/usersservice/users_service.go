package usersservice

import (
	"context"
	"log"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/usersservice/pb"
	"google.golang.org/grpc"
)

func CreateUser(sessionId string) (string, error) {
	conn, err := grpc.Dial(os.Getenv("USERS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to users service: %v", err)
		return "", err
	}
	defer conn.Close()

	client := pb.NewUsersServiceClient(conn)
	request := &pb.CreateUserRequest{SessionId: sessionId}

	resp, err := client.CreateUser(context.Background(), request)

	if err != nil {
		return "", err
	}

	return resp.UserId, nil
}

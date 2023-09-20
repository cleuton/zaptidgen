package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/cleuton/zaptidgen/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := "localhost:8888"
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		panic("did not connect")
	}
	defer conn.Close()
	c := pb.NewIdGenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	idRequest := pb.IdRequest{}

	idResponse, err := c.Gen(ctx, &idRequest)

	fmt.Println("ID: ", idResponse.Id)
}

package main

import (
	"context"
	"fmt"
	"marmota/api"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := api.NewNodeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.UpdateNode(ctx, &api.UpdateNodeReq{})

	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

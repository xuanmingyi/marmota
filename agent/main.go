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

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	for {
		r, err := c.UpdateNode(ctx, &api.UpdateNodeReq{
			Item: &api.NodeItem{
				Id:        1,
				Uuid:      "abcdefghijklmn",
				Name:      "xxxxxxxxx",
				Metadata:  "{xxx:bbb, ccc:dddd}",
				Desc:      "ssssssssss",
				CreateAt:  11111111,
				UpdatedAt: 22222222,
			},
		})

		if err != nil {
			panic(err)
		}
		fmt.Println(r)
		time.Sleep(time.Second)
	}
}

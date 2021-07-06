package main

import (
	"fmt"
	"marmota/api"
	"time"
)

func main() {
	conf, err := ReadConfig("configs.yaml")
	if err != nil {
		panic(err)
	}

	client, err := NewClient(conf)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			r, _ := client.UpdateNode(&api.UpdateNodeReq{
				Item: &api.NodeItem{
					Uuid: conf.Agent.Uuid,
					Name: conf.Agent.Name,
					Desc: conf.Agent.Desc,
				},
			})
			fmt.Println(r)
			time.Sleep(time.Duration(conf.Agent.Interval) * time.Second)
		}
	}()

	select {}
}

package main

import (
	"github.com/coreos/etcd/client"
	"fmt"
	"golang.org/x/net/context"
)

func main()  {
	machines := []string{"http://172.17.0.2:2379"}
	cfg := client.Config{
		Endpoints:machines,
		Transport:client.DefaultTransport,
	}

	cli, err := client.New(cfg)
	if err != nil {
		fmt.Println(err)
	}

	kapi := client.NewKeysAPI(cli)
	fmt.Println("setting '/foo', key")
	rsp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("set is done %q", rsp)
	}

}
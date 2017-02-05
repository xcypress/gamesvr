package main

import (
	etcdclient "github.com/coreos/etcd/client"
	"fmt"
	"golang.org/x/net/context"
)

const 	DEFAULT_NAME_FILE    = "/backends/redis"


func main()  {
	machines := []string{"http://172.17.0.2:2379"}
	cfg := etcdclient.Config{
		Endpoints:machines,
		Transport:etcdclient.DefaultTransport,
	}
    client, err := etcdclient.New(cfg)
    if err != nil {
        fmt.Println(err)
    }

    kApi := etcdclient.NewKeysAPI(client)
    rsp, err := kApi.Get(context.Background(), DEFAULT_NAME_FILE, nil)
    if err != nil {
        fmt.Println(err)
    }
    for _, node := range rsp.Node.Nodes {
        fmt.Println(node.Key)
        fmt.Println(node.Value)
    }
}
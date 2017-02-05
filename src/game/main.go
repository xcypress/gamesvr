package main

import (
	"github.com/coreos/etcd/client"
	"fmt"
	"golang.org/x/net/context"
)

const   DEFAULT_SERVIE_KEY   = "/redis"


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

    kApi := client.NewKeysAPI(cli)
    //rsp, err := kApi.Get(context.Background(), DEFAULT_NAME_FILE, nil)
    //if err != nil {
    //    fmt.Println(err)
    //}
    //
    //if rsp.Node.Dir {
    //    fmt.Println("names is not a file")
    //}
    //
    //
    //fmt.Println(rsp.Node.Value)
    //fmt.Println(strings.Split(rsp.Node.Value, "\n"))

    rsp, err := kApi.Get(context.Background(), DEFAULT_SERVIE_KEY, nil)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("continue")

    fmt.Println(rsp.Node.Value)
}
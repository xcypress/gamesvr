package main

import (
    "service"
    "fmt"
)

func main()  {
    serviceMgr := new(service.ServiceMgr)
    serviceMgr.Init()
    fmt.Println("main start")
    for {
        select {
        case node := <-serviceMgr.AddServiceMQ:
            serviceMgr.AddService(node.Key, node.Conn)
        case key := <-serviceMgr.RemoveServiceMQ:
            serviceMgr.RemoveService(key)
        }
    }
}
package main

import "service"

func main()  {
    serviceMgr := new(service.ServiceMgr)
    for {
        select {
        case node := <-serviceMgr.AddServiceMQ:
            serviceMgr.AddService(node.Key, node.Conn)
        case key := <-serviceMgr.RemoveServiceMQ:
            serviceMgr.RemoveService(key)
        }
    }
}
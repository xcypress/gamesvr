package service

import  (
    "net"
    "github.com/coreos/etcd/client"
    "context"
    "log"
    "os"
    "path/filepath"
    "strings"
    "fmt"
)

const (
    DEFAULT_SERVICE_PATH = "/backends"
    DEFAULT_NAME_FILE = "backends/names"
    DEFAULT_ETCD_HOST = "172.17.0.2:2379"
)

type Node struct {
    Key  string
    Conn *net.TCPConn
}
type service struct {
    nodes   []*Node
    idx     uint32
}
type ServiceMgr struct {
    services map[string]*service
    known_names     map[string]bool
    AddServiceMQ    chan Node
    RemoveServiceMQ        chan string
    etcdClient      client.Client
}

func (sm *ServiceMgr) init() {

    machines := []string{DEFAULT_ETCD_HOST}
    if env := os.Getenv("ETCD_HOST"); env != "" {
        machines = strings.Split(env, ";")
    }

    // init etcd client
    cfg := client.Config{
        Endpoints: machines,
        Transport: client.DefaultTransport,
    }

    cli, err := client.New(cfg)
    if err != nil {
        log.Panic(err)
    }
    sm.etcdClient = cli

    sm.services = make(map[string]*service)
    sm.known_names = make(map[string]bool)

    names := sm.loadNames()

    for _, name := range names {
        sm.known_names[DEFAULT_SERVICE_PATH + "/" + strings.TrimSpace(name)] = true
    }

    sm.connectAll(DEFAULT_SERVICE_PATH)

}
func (sm *ServiceMgr) watcher() {
    kApi := client.NewKeysAPI(sm.etcdClient)
    watcher := kApi.Watcher(DEFAULT_SERVICE_PATH, &client.WatcherOptions{Recursive:true})

    for {
        rsp, err := watcher.Next(context.Background())
        if err != nil {
            log.Println(err)
            continue
        }

        if rsp.Node.Dir {
            continue
        }

        switch rsp.Action {
        case "set", "create", "update", "compareAndSwap":
            conn, err := net.DialTCP("tcp", nil, rsp.Node.Value)
            if err == nil {
                sm.AddServiceMQ <- Node{rsp.Node.Key, conn}
            } else {
                log.Println(err)
                log.Println("can not connect ",rsp.Node.Key, rsp.Node.Value)
            }

        case "delete":
            sm.RemoveServiceMQ <- rsp.Node.Key
        
        }
    }
}

func (sm *ServiceMgr) loadNames() []string {
    kApi := client.NewKeysAPI(sm.etcdClient)
    log.Println("reading names :", DEFAULT_NAME_FILE)
    rsp, err := kApi.Get(context.Background(), DEFAULT_NAME_FILE, nil)
    if err != nil {
        fmt.Println(err)
        return nil
    }

    if rsp.Node.Dir {
        log.Println("names is not a file")
    }

    return strings.Split(rsp.Node.Value, "/n")
}

func (sm *ServiceMgr) connectAll(dir string) {
    kApi := client.NewKeysAPI(sm.etcdClient)
    log.Println("connecting services under:", dir)
    rsp, err := kApi.Get(context.Background(), dir, &client.GetOptions{Recursive: true})
    if err != nil {
        log.Println(err)
        return
    }

    for _, node := range rsp.Node.Nodes {
        if node.Dir {
            for _, service := range node.Nodes {
                service_name := filepath.Dir(service.Key)
                if !sm.known_names[service_name] {
                    return
                }

                conn, err := net.DialTCP("tcp", nil, service.Value)
                if err == nil {
                    sm.AddService(service.Key, conn)
                    log.Println("connect service:" + service.Value)
                }
            }
        }
    }
    log.Println("services add complete")
    go sm.watcher()
}

func (sm *ServiceMgr) AddService(key string, conn *net.TCPConn) {
    serviceName := filepath.Dir(key)
    if !sm.known_names[serviceName] {
        return
    }

    if sm.services[serviceName] == nil {
        sm.services[serviceName] = &service{}
    }

    service := sm.services[serviceName]
    node := &Node{
        Key: key,
        Conn: conn,
    }
    service.nodes = append(service.nodes, node)
}

func (sm *ServiceMgr) RemoveService(key string) {
    serviceName := filepath.Dir(key)
    if !sm.known_names[serviceName] {
        return
    }
    service := sm.services[serviceName]
    if service == nil {
        log.Println("no such service:", serviceName)
        return
    }

    for idx := range service.nodes {
        if service.nodes[idx].Key == key {
            service.nodes[idx].Conn.Close()
            service.nodes = append(service.nodes[:idx], service.nodes[idx+1]...)
            log.Println("service remove:", key)
            return
        }
    }
}


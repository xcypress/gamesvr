https://hub.docker.com/r/gliderlabs/registrator/
>docker pull gliderlabs/registrator

>docker run -d  --name=registrator --net=host  --volume=/var/run/docker.sock:/tmp/docker.sock  gliderlabs/registrator:latest etcd://172.17.0.2:4001/backends


etcd  https://hub.docker.com/r/elcolio/etcd/
>docker pull elcolio/etcd
>docker run -d -p 2379:2379 -p 4001:4001 --name=etcd1 elcolio/etcd  2379 http 4001 peer

etcd-browser https://hub.docker.com/r/denverdino/etcdbrowser/
> docker pull denverdino/etcdbrowser

>docker run --rm --name etcd-browser -p 0.0.0.0:8000:8000 --env ETCD_HOST=172.17.0.2 --env AUTH_PASS=doe -t -i denverdino/etcdbrowser

http://cizixs.com/2016/08/02/intro-to-etcd

user govendor
gopath -> src -> game ->vendor
go get -u -v github.com/kardianos/govendor 
govendor init
govendor fetch golang.org/x/net/context


    init     Create the "vendor" folder and the "vendor.json" file.

    list     List and filter existing dependencies and packages.

    add      Add packages from $GOPATH.

    update   Update packages from $GOPATH.

    remove   Remove packages from the vendor folder.

    status   Lists any packages missing, out-of-date, or modified locally.

    fetch    Add new or update vendor folder packages from remote repository.

    sync     Pull packages into vendor folder from remote repository with revisions

                 from vendor.json file.

    migrate  Move packages from a legacy tool to the vendor folder with metadata.

    get      Like "go get" but copies dependencies into a "vendor" folder.

    license  List discovered licenses for the given status or import paths.

    shell    Run a "shell" to make multiple sub-commands more efficient for large

                 projects.

#CAP理论
> * 一致性
> * 可用性
> * 分区容忍性

服务与发现满足CP
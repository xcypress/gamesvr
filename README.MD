register https://hub.docker.com/r/jwilder/docker-register/ 
>docker pull jwilder/docker-register

>docker run -d -e ETCD_HOST=172.17.0.2:4001 -v /var/run/docker.sock:/var/run/docker.sock -t jwilder/docker-register

etcd  https://hub.docker.com/r/elcolio/etcd/
>docker pull elcolio/etcd
>docker run -d -p 2379:2379 -p 4001:4001 -name=etcd1 elcolio/ectd  2379 http 4001 peer

etcd-browser https://hub.docker.com/r/denverdino/etcdbrowser/
> docker pull denverdino/etcdbrowser

>docker run --rm --name etcd-browser -p 0.0.0.0:8000:8000 --env ETCD_HOST=172.17.0.2 --env AUTH_PASS=doe -t -i denverdino/etcdbrowser

http://cizixs.com/2016/08/02/intro-to-etcd

user govendor
gopath -> src -> game ->vendor
go get -u -v github.com/kardianos/govendor 
govendor init
govendor fetch golang.org/x/net/context
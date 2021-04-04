module github.com/go-alive/examples

go 1.13

replace github.com/go-alive/go-micro => /Users/unlikezhang/go/src/github.com/go-alive/go-micro

replace github.com/go-alive/micro => /Users/unlikezhang/go/src/github.com/go-alive/micro

replace github.com/go-alive/cli => /Users/unlikezhang/go/src/github.com/go-alive/cli

//replace github.com/go-alive/go-plugins => /Users/unlikezhang/go/src/github.com/go-alive/go-plugins

replace github.com/go-alive/go-plugins/registry/zookeeper => /Users/unlikezhang/go/src/github.com/go-alive/go-plugins/registry/zookeeper

replace k8s.io/api => k8s.io/api v0.0.0-20190708174958-539a33f6e817

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190708180123-608cd7da68f7

replace k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible

replace k8s.io/component-base => k8s.io/component-base v0.0.0-20190708175518-244289f83105

replace google.golang.org/grpc => google.golang.org/grpc v1.24.0

require (
	github.com/99designs/gqlgen v0.10.1
	github.com/astaxie/beego v1.12.0
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/emicklei/go-restful v2.11.1+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/go-alive/cli v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-alive/go-micro v1.0.0
	github.com/go-alive/go-plugins/registry/zookeeper v0.0.0-00010101000000-000000000000
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.4.1
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/pborman/uuid v1.2.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/vektah/gqlparser v1.2.0
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/grpc v1.26.0
)

package app

import (
	"flag"
	"sync"
)

type Param struct {
	HttpPort        *int    // http端口
	GrpcPort        *int    // grpc端口
	RouteServerAddr *string // 路由服务地址
	RouteServerPort *int    // 路由服务端口
	HttpNative      *bool   // 是否http原生server
}

var instance *Param
var once sync.Once

// GetInstance 获取实例
func GetInstance() *Param {
	once.Do(func() {
		instance = &Param{}
	})
	return instance
}

func (p *Param) Parse(flag *flag.FlagSet) {
	var routeServerAddr string
	var httpPort int
	var grpcPort int
	var routeServerPort int
	var httpNativeServer bool

	flag.StringVar(&routeServerAddr, "route.server.address", "localhost", "路由服务地址")
	flag.IntVar(&httpPort, "http.port", 7654, "http端口")
	flag.IntVar(&grpcPort, "grpc.port", 6543, "grpc端口")
	flag.IntVar(&routeServerPort, "route.server.port", 5432, "路由服务端口")
	flag.BoolVar(&httpNativeServer, "http.native.server", false, "是否采用go http原生server")

	p.RouteServerAddr = &routeServerAddr
	p.HttpPort = &httpPort
	p.RouteServerPort = &routeServerPort
	p.GrpcPort = &grpcPort
	p.HttpNative = &httpNativeServer
}

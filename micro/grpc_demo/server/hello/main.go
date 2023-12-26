package main

import (
	"context"
	"fmt"
	"ginDemo/micro/grpc_demo/server/hello/helloService"
	"google.golang.org/grpc"
	"net"
)

//rpc远程调用的接口,需要实现hello.proto中定义的Hello接口,以及里面的方法
//1.定义远程调用的结构体和方法,这个结构体需要实现HelloServer的接口

type Hello struct{}

func (h Hello) SayHello(c context.Context, req *helloService.HelloReq) (*helloService.HelloRes, error) {
	fmt.Println(req)
	return &helloService.HelloRes{
		Message: "你好" + req.Name,
	}, nil
}

func main() {
	//1. 初始一个 grpc 对象
	grpcServer := grpc.NewServer()
	//2. 注册服务
	//helloService.RegisterHelloServer(grpcServer, &Hello{})
	// &Hello{}和 new(Hello)相同
	helloService.RegisterHelloServer(grpcServer, &Hello{})
	//3. 设置监听， 指定 IP、port
	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}
	// 4退出关闭监听
	defer listen.Close()
	// 5. 启动服务
	grpcServer.Serve(listen)

}

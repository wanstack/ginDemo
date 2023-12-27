package main

import (
	"context"
	"fmt"
	"ginDemo/micro/grpc_demo/server/hello/helloService"
	"github.com/hashicorp/consul/api"
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

	// --------------- consul 服务相关 ----------
	//1、初始化consul配置
	consulConfig := api.DefaultConfig()
	//设置consul服务器地址: 默认127.0.0.1:8500, 如果consul部署到其它服务器上,则填写其它服务器地址
	//consulConfig.Address = "127.0.0.1:8500"
	consulConfig.Address = "192.168.3.10:8500"
	//2、获取consul操作对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println(err)
	}
	// 3、配置注册服务的参数
	agentService := api.AgentServiceRegistration{
		ID:      "1",              // 服务id,顺序填写即可
		Tags:    []string{"test"}, // tag标签
		Name:    "HelloService",   //服务名称, 注册到服务发现(consul)的K
		Port:    8089,             // 端口号: 需要与下面的监听， 指定 IP、port一致
		Address: "192.168.3.78",   // 当前微服务部署地址: 结合Port在consul设置为V: 需要与下面的监听， 指定 IP、port一致
		Check: &api.AgentServiceCheck{ //健康检测
			TCP:      "192.168.3.78:8089", //前微服务部署地址,端口 : 需要与下面的监听， 指定 IP、port一致
			Timeout:  "5s",                // 超时时间
			Interval: "30s",               // 循环检测间隔时间
		},
	}
	//4、注册服务到consul上
	consulClient.Agent().ServiceRegister(&agentService)

	// -------------- grpc 相关---------------
	//1. 初始一个 grpc 对象
	grpcServer := grpc.NewServer()
	//2. 注册服务
	//helloService.RegisterHelloServer(grpcServer, &Hello{})
	// &Hello{}和 new(Hello)相同
	helloService.RegisterHelloServer(grpcServer, &Hello{})
	//3. 设置监听， 指定 IP、port
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		panic(err)
	}
	// 4退出关闭监听
	defer listen.Close()
	// 5. 启动服务
	grpcServer.Serve(listen)

}

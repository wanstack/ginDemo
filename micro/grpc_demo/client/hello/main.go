package main

import (
	"context"
	"fmt"
	"ginDemo/micro/grpc_demo/server/hello/helloService"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

// 1、连接服务器
/*
   credentials.NewClientTLSFromFile ：从输入的证书文件中为客户端构造TLS凭证。
   grpc.WithTransportCredentials ：配置连接级别的安全凭证（例如，TLS/SSL），返回一个
   DialOption，用于连接服务器。
*/
func main() {

	//----------------------------consul相关---------------------------
	//初始化consul配置, 客户端服务器需要一致
	consulConfig := api.DefaultConfig()
	//设置consul服务器地址: 默认127.0.0.1:8500, 如果consul部署到其它服务器上,则填写其它服务器地址
	//consulConfig.Address = "127.0.0.1:8500"
	consulConfig.Address = "192.168.3.10:8500"
	//2、获取consul操作对象
	consulClient, _ := api.NewClient(consulConfig) //目前先屏蔽error,也可以获取error进行错误处理
	//3、获取consul服务发现地址,返回的ServiceEntry是一个结构体数组
	//参数说明:service：服务名称,服务端设置的那个Name, tag:标签,服务端设置的那个Tags,, passingOnly bool, q: 参数
	serviceEntry, _, _ := consulClient.Health().Service("HelloService", "test", false, nil)
	//打印地址
	fmt.Printf("%#v\n", serviceEntry)
	fmt.Println(serviceEntry[0].Service.Address)
	fmt.Println(serviceEntry[0].Service.Port)
	//拼接地址
	//strconv.Itoa: int转string型
	address := serviceEntry[0].Service.Address + ":" + strconv.Itoa(serviceEntry[0].Service.Port)

	// --------------- grpc 相关--------

	grpcClient, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	//2、注册客户端
	client := helloService.NewHelloClient(grpcClient)

	//3、调用服务端函数, 实现HelloClient接口:SayHello()
	res, err := client.SayHello(context.Background(), &helloService.HelloReq{
		Name: "zhangsan",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res.Message)

}

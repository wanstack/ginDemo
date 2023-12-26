package main

import (
	"context"
	"fmt"
	"ginDemo/micro/grpc_demo/server/hello/helloService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 1、连接服务器
/*
   credentials.NewClientTLSFromFile ：从输入的证书文件中为客户端构造TLS凭证。
   grpc.WithTransportCredentials ：配置连接级别的安全凭证（例如，TLS/SSL），返回一个
   DialOption，用于连接服务器。
*/
func main() {
	grpcClient, err := grpc.Dial("127.0.0.1:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

package main

import (
	"context"
	"fmt"
	"ginDemo/micro/grpc_demo/server/goods/proto/goodsService"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

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
	//serviceEntry, _, _ := consulClient.Health().Service("HelloService", "test", false, nil)
	//打印地址
	//fmt.Println(serviceEntry[0].Service.Address)
	//fmt.Println(serviceEntry[0].Service.Port)
	//拼接地址
	//strconv.Itoa: int转string型
	//address := serviceEntry[0].Service.Address + ":" + strconv.Itoa(serviceEntry[0].Service.Port)

	//----------------------------hello微服务相关------------------------------
	// 1、连接服务器
	/*
	   credentials.NewClientTLSFromFile ：从输入的证书文件中为客户端构造TLS凭证。
	   grpc.WithTransportCredentials ：配置连接级别的安全凭证（例如，TLS/SSL），返回一个
	   DialOption，用于连接服务器。
	*/
	//把上面拼接的地址放入下面
	//grpcClient, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//    fmt.Println(err)
	//}

	//2、注册客户端
	//client := helloService.NewHelloClient(grpcClient)
	//3、调用服务端函数, 实现HelloClient接口:SayHello()
	/*
	   // HelloClient is the client API for Hello service.
	   //
	   // For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
	   type HelloClient interface {
	       // 通过rpc来指定远程调用的方法:
	       // SayHello方法, 这个方法里面实现对传入的参数HelloReq, 以及返回的参数HelloRes进行约束
	       SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error)
	   }
	*/
	//res, err1 := client.SayHello(context.Background(), &helloService.HelloReq{
	//    Name: "张三",
	//})
	//if err1 != nil {
	//    fmt.Printf("调用服务端代码失败: %s", err1)
	//    return
	//}
	//
	//fmt.Printf("%#v\r\n", res)
	//fmt.Printf("调用成功: %s", res.Message)

	// ----------------------------goods微服务相关--------------------------

	//3、获取consul服务发现地址,返回的ServiceEntry是一个结构体数组
	//参数说明:service：服务名称,服务端设置的那个Name, tag:标签,服务端设置的那个Tags,, passingOnly bool, q: 参数
	serviceGoodsEntry, _, _ := consulClient.Health().Service("GoodsService", "test", false, nil)
	//打印地址
	fmt.Println(serviceGoodsEntry[0].Service.Address)
	fmt.Println(serviceGoodsEntry[0].Service.Port)
	//拼接地址
	//strconv.Itoa: int转string型
	addressGoods := serviceGoodsEntry[0].Service.Address + ":" + strconv.Itoa(serviceGoodsEntry[0].Service.Port)

	// 1、连接服务器
	/*
	   credentials.NewClientTLSFromFile ：从输入的证书文件中为客户端构造TLS凭证。
	   grpc.WithTransportCredentials ：配置连接级别的安全凭证（例如，TLS/SSL），返回一个
	   DialOption，用于连接服务器。
	*/

	// 连接
	grpcClient, err := grpc.Dial(addressGoods, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	// 注册客户端
	client := goodsService.NewGoodsClient(grpcClient)

	// 调用服务端方法
	res, err := client.AddGoods(context.Background(), &goodsService.AddGoodsReq{
		Goods: &goodsService.GoodsModel{
			Title:   "测试商品",
			Price:   20,
			Content: "测试商品内容",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res.Message)

	// 获取商品数据
	res2, _ := client.GetGoods(context.Background(), &goodsService.GetGoodsReq{})
	fmt.Printf("%#v\n", res2.GoodsList)

	for i := 0; i < len(res2.GoodsList); i++ {
		fmt.Printf("%#v\r\n", res2.GoodsList[i])
	}
}

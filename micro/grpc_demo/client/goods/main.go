package main

import (
	"context"
	"fmt"
	"ginDemo/micro/grpc_demo/server/goods/proto/goodsService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接
	grpcClient, err := grpc.Dial("127.0.0.1:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

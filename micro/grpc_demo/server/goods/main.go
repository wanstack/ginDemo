package main

import (
	"context"
	"fmt"
	"ginDemo/micro/grpc_demo/server/goods/proto/goodsService"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

//rpc远程调用的接口,需要实现goods.proto中定义的Goods服务接口,以及里面的方法
//1.定义远程调用的结构体和方法,这个结构体需要实现GoodsServer的接口

type Goods struct{}

//GoodsServer方法参考goods.pb.go中的接口
/*
type GoodsServer interface {
	//通过rpc来指定远程调用的方法:
	//AddGoods方法, 这个方法里面实现对传入的参数AddGoodsReq, 以及返回的参数AddGoodsRes进行约束
	AddGoods(context.Context, *AddGoodsReq) (*AddGoodsRes, error)
}
*/

func (g Goods) AddGoods(c context.Context, req *goodsService.AddGoodsReq) (*goodsService.AddGoodsRes, error) {
	fmt.Println("req:", req)
	return &goodsService.AddGoodsRes{
		Message: "增加成功" + req.Goods.Title,
		Success: true,
	}, nil
}

func (g Goods) GetGoods(c context.Context, req *goodsService.GetGoodsReq) (*goodsService.GetGoodsRes, error) {
	var tempList []*goodsService.GoodsModel //定义返回的商品列表切片
	//模拟从数据库中获取商品的请求,循环结果,把商品相关数据放入tempList切片中
	for i := 0; i < 10; i++ {
		tempList = append(tempList, &goodsService.GoodsModel{
			Title:   "商品" + strconv.Itoa(i), // strconv.Itoa(i): 整型转字符串类型
			Price:   float64(i),             //float64(i): 强制转换整型为浮点型
			Content: "测试商品内容" + strconv.Itoa(i),
		})
	}
	return &goodsService.GetGoodsRes{
		GoodsList: tempList,
	}, nil
}

func main() {
	// 初始化一个rpc对象
	grpcServer := grpc.NewServer()
	// 注册服务端
	goodsService.RegisterGoodsServer(grpcServer, new(Goods))
	// 设置监听地址和端口
	listen, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println(err)
	}
	// 退出关闭监听
	defer listen.Close()
	// 启动服务
	grpcServer.Serve(listen)

}

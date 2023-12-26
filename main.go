package main

import (
	"fmt"
	"ginDemo/models"
	"ginDemo/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
)

func init() {
	// 创建表
	models.CreateTable()
}

func main() {
	r := gin.Default()

	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatal("redis配置文件错误")
	}
	ip := config.Section("redis").Key("ip").String()
	port := config.Section("redis").Key("port").String()
	password := config.Section("redis").Key("password").String()
	//database := config.Section("redis").Key("database").Ints()

	// 使用redis保存session,利用session机制保持登录
	//参数说明:
	//自第1个参数-redis最大的空闲连接数
	//第2个参数-数通信协议tcp或者udp
	//第3个参数-redis地址,格式，host:port 第4个参数-redis密码
	//第5个参数-session加密密钥
	store, _ := redis.NewStore(10, "tcp", fmt.Sprintf("%v:%v", ip, port), password, []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//加载templates中所有模板文件, 使用不同目录下名称相同的模板,注意:一定要放在配置路由之前才得行
	//如果模板在多级目录里面的话需要这样配置 r.LoadHTMLGlob("templates/**/**/*") /** 表示目录
	//LoadHTMLGlob只能加载同一层级的文件
	//比如说使用router.LoadHTMLFile("/templates/**/*")，就只能加载/templates/admin/或者/templates/order/下面的文件
	//解决办法就是通过filepath.Walk来搜索/templates下的以.html结尾的文件，把这些html文件都加载一个数组中，然后用LoadHTMLFiles加载
	//var files []string
	//filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
	//    if strings.HasSuffix(path, ".html") {
	//        files = append(files, path)
	//    }
	//    return nil
	//})
	//r.LoadHTMLFiles(files...)
	r.LoadHTMLGlob("templates/**/**/*")

	//配置静态web目录 第一个参数表示路由,第二个参数表示映射的目录
	r.Static("/static", "./static")
	//分组路由文件
	routers.AdminRoutersInit(r)
	//routers.ApiRoutersInit(r)
	//routers.DefaultRoutersInit(r)

	r.Run(":8080")

}

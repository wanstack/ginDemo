package middlewares

import (
	"encoding/json"
	"fmt"
	"ginDemo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"net/http"
	"os"
	"strings"
)

// InitAdminAuthMiddleware 中间件: 作用: 在执行路由之前或者之后进行相关逻辑判断
func InitAdminAuthMiddleware(c *gin.Context) {
	// 权限判断: 没有登录的用户不能进入后台管理中心
	//1、获取Url访问的地址
	//当地址后面带参数时:,如: admin/role?title=test,需要处理
	//strings.Split(c.Request.URL.String(), "?"): 把c.Request.URL.String()请求地址按照?分割成切片
	pathname := strings.Split(c.Request.URL.String(), "?")[0] // /admin/role

	//2、获取Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")

	//3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行
	//session.Get获取返回的结果是一个空接口类型,所以需要进行类型断言: 判断userinfo是不是一个string
	userinfoStr, ok := userinfo.(string) //类型断言
	if ok {                              // 说明用户名存在
		var u []models.Manager
		//把获取到的用户信息转换结构体
		json.Unmarshal([]byte(userinfoStr), &u)
		if !(len(u) > 0 && u[0].Username != "") {
			if pathname != "/admin/login" {
				//跳转到登录页面
				//c.Redirect(302, "/admin/login")
				c.JSON(302, "请到/admin/login页面登录")
			}
		} else { // 用户存在，判断权限
			//获取当前访问的URL对应的权限id,判断权限id是否在角色对应的权限中
			// strings.Replace 字符串替换, 此为去掉/admin/ 字符串
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			//排除权限判断:不是超级管理员并且不在相关权限内
			if u[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
				//判断用户权限:当前用户权限是否可以访问url地址
				//获取当前角色拥有的权限,并把权限id放在一个map对象中
				var roleAccess []models.RoleAccess
				models.DB.Where("role_id = ?", u[0].RoleId).Find(&roleAccess)
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId
				}
				//实例化access
				access := models.Access{}
				//查询权限id
				models.DB.Where("url = ? ", urlPath).Find(&access)
				//判断权限id是否在角色对应的权限中
				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(http.StatusOK, "没有权限")
					c.Abort() // 终止程序
				}
			} else {
				if pathname != "/admin/login" {
					c.JSON(302, "请到/admin/login页面登录")
				}
			}
		}
	}
}

//排除权限判断的方法
func excludeAuthPath(urlPath string) bool {
	//加载配置文件
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	//获取需要排除的地址
	excludeAuthPath := cfg.Section("NoAuthPath").Key("excludeAuthPath").String()
	//拆分字符串成为一个切片
	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	//判断传入的地址是否在排除地址内
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}

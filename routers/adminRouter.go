package routers

import (
	"ginDemo/controller/admin"
	"ginDemo/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//路由分组: 配置全局中间件:middlewares.InitMiddleware
	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/", admin.IndexController{}.Index)          // 实例化控制器,并访问其中方法
		adminRouters.GET("/welcome", admin.IndexController{}.Welcome) // 实例化控制器,并访问其中方法

		// login 登录后才可以访问/admin 下的其他资源
		adminRouters.GET("/login", admin.LoginController{}.Login)
		adminRouters.GET("/logout", admin.LoginController{}.Logout)

		// 角色
		adminRouters.GET("/role", admin.RoleController{}.Index)
		adminRouters.POST("/role", admin.RoleController{}.Create)
		adminRouters.PUT("/role", admin.RoleController{}.Update)
		adminRouters.DELETE("/role", admin.RoleController{}.Delete)

		// 管理员
		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.POST("/manager", admin.ManagerController{}.Create)
		adminRouters.PUT("/manager", admin.ManagerController{}.Update)
		adminRouters.DELETE("/manager", admin.ManagerController{}.Delete)

		// 权限路由
		adminRouters.GET("/access", admin.AccessController{}.Index)
		adminRouters.POST("/access", admin.AccessController{}.Create)
		adminRouters.PUT("/access", admin.AccessController{}.Update)
		adminRouters.DELETE("/access", admin.AccessController{}.Delete)

		// 授权
		adminRouters.POST("/role/auth", admin.RoleController{}.Auth)
		adminRouters.PUT("/role/auth", admin.RoleController{}.Auth)

	}
}

package admin

import (
	"encoding/json"
	"ginDemo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

// Login 执行登录操作
func (con LoginController) Login(c *gin.Context) {
	//获取用户名以及密码
	username := c.Query("username")
	password := c.Query("password")

	//查询数据库,判断用户以及密码是否正确
	var userinfo []models.Manager
	password = models.Md5(password)
	// 获取数据库中受影响的行数
	affected := models.DB.
		Where("username = ? and password = ? ", username, password).Find(&userinfo).RowsAffected
	if affected > 0 {
		//执行登录,保存用户信息,执行跳转操作
		session := sessions.Default(c)
		//注意: session.Set没法保存结构体对应的切片,所以需要把结构体转换成json字符串
		userinfoSlice, _ := json.Marshal(userinfo)
		session.Set("userinfo", string(userinfoSlice))
		session.Save()
		//c.HTML(http.StatusOK, "/admin/", "")
		con.Success(c, "登录成功,跳转到/admin/ 页面")
	} else {
		con.Error(c, "登录失败,请圆润的离开重新/admin/login")
	}
}

func (con LoginController) Logout(c *gin.Context) {
	//1.销毁session中用户信息
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.Success(c, "登出成功,欢迎下次光临")
}

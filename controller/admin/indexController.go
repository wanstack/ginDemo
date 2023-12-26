package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (con IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index/index.html", gin.H{
		"username": "超级管理员",
	})
}

func (con IndexController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/welcome.html", gin.H{})
}

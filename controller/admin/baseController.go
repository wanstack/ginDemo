package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

// Success 公共成功方法
func (BaseController) Success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

// Error 公共失败方法
func (BaseController) Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

package admin

import (
	"ginDemo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type RoleController struct {
	BaseController
}

// Index 获取角色列表
func (RoleController) Index(c *gin.Context) {
	// 定义一个角色切片
	var roleList []models.Role
	// 获取所有角色
	models.DB.Find(&roleList)
	c.JSON(http.StatusOK, gin.H{
		"roleList": roleList,
	})

}

// Create 新增角色
func (con RoleController) Create(c *gin.Context) {
	// 获取表单提交的数据
	title := strings.Trim(c.Query("title"), " ") // 去掉字符串两边空格
	description := strings.Trim(c.Query("description"), " ")

	if title == "" {
		con.Error(c, "角色不能为空")
		return
	}
	err := models.DB.Create(&models.Role{
		Title:       title,
		Description: description,
		Status:      1,
		AddTime:     int(time.Now().Unix()),
	}).Error
	if err != nil {
		con.Error(c, "新增角色失败")
		return
	}
	con.Success(c, "新增角色成功")
}

// Update 编辑角色
func (con RoleController) Update(c *gin.Context) {
	// 获取提交表单数据
	id, err := models.StringToInt(strings.Trim(c.Query("id"), " "))
	if err != nil {
		con.Error(c, "传入id错误")
		return
	}
	title := strings.Trim(c.Query("title"), " ")
	description := strings.Trim(c.Query("description"), " ")
	//判断角色名称是否为空
	if title == "" {
		con.Error(c, "角色名称不能为空")
		return
	}
	// 查询到id后更新操作
	err = models.DB.Model(&models.Role{}).Where("id = ?", id).Updates(models.Role{
		Id:          id,
		Title:       title,
		Description: description,
		AddTime:     int(models.GetUnix()),
	}).Error
	if err != nil {
		con.Error(c, "修改数据失败")
		return
	}
	con.Success(c, "修改数据成功")

}

// Delete 删除角色
func (con RoleController) Delete(c *gin.Context) {
	id, err := models.StringToInt(strings.Trim(c.Query("id"), " "))
	if err != nil {
		con.Error(c, "传入数据错误")
		return
	}
	err = models.DB.Delete(&models.Role{}, []int{id}).Error
	if err != nil {
		con.Error(c, "删除失败")
	}
	con.Success(c, "删除成功")

}

// Auth 授权
func (con RoleController) Auth(c *gin.Context) {
	// 获取roleId
	roleId, err := models.StringToInt(c.Query("role_id"))
	if err != nil {
		con.Error(c, "role不存在")
		return
	}
	//获取表单提交的权限id切片
	accessIds := c.QueryArray("access_node[]")
	//清空当前角色权限id
	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id = ?", roleId).Delete(&roleAccess)

	//循环遍历accessIds,增加当前角色对应的权限
	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		accessId, _ := models.StringToInt(v)
		roleAccess.AccessId = accessId
		models.DB.Create(&roleAccess)
	}

	con.Success(c, "角色授权成功")
}

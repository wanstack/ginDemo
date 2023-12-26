package admin

import (
	"ginDemo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ManagerController struct {
	BaseController
}

// Index 获取管理员列表以及对应的角色
func (con ManagerController) Index(c *gin.Context) {
	var managerList []models.Manager
	models.DB.Preload("Role").Find(&managerList)
	c.JSON(http.StatusOK, gin.H{
		"managerList": managerList,
	})
}

// Create 添加管理员
func (con ManagerController) Create(c *gin.Context) {
	// 获取角色id
	RoleId, err := models.StringToInt(c.Query("role_id"))
	if err != nil {
		con.Error(c, "角色id不合法")
		return
	}

	// 获取提交的管理员信息
	username := strings.Trim(c.Query("username"), " ")
	password := strings.Trim(c.Query("password"), " ")
	email := strings.Trim(c.Query("email"), " ")
	mobile := strings.Trim(c.Query("mobile"), " ")

	// 判断管理员是否存在
	var managerList []models.Manager
	affected := models.DB.Where("username = ?", username).Find(&managerList).RowsAffected
	if affected != 0 {
		con.Error(c, "管理员已经存在")
	}
	// 创建
	err = models.DB.Create(&models.Manager{
		Username: username,
		Password: models.Md5(password),
		Email:    email,
		Mobile:   mobile,
		AddTime:  int(models.GetUnix()),
		RoleId:   RoleId,
		Status:   1,
	}).Error
	if err != nil {
		con.Error(c, "添加管理员失败")
		return
	}
	con.Success(c, "添加管理员成功")
}

// Update 更新管理员
func (con ManagerController) Update(c *gin.Context) {
	// 获取管理员id
	id, err := models.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, "管理员不存在")
		return
	}

	// 获取角色id
	roleId, err := models.StringToInt(c.Query("role_id"))
	if err != nil {
		con.Error(c, "角色不存在")
		return
	}
	// 获取管理员需要修改的信息
	username := strings.Trim(c.Query("username"), " ")
	password := strings.Trim(c.Query("password"), " ")
	mobile := strings.Trim(c.Query("mobile"), " ")
	email := strings.Trim(c.Query("email"), " ")
	// 修改
	err = models.DB.Where("id = ?", id).Updates(&models.Manager{
		Username: username,
		Password: models.Md5(password),
		Mobile:   mobile,
		Email:    email,
		RoleId:   roleId,
	}).Error
	if err != nil {
		con.Error(c, "修改失败")
	}
	con.Success(c, "修改成功")
}

// Delete 删除管理员
func (con ManagerController) Delete(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, "管理员不存在")
		return
	}

	//查询管理员是否存在
	var managerList []models.Manager
	affected := models.DB.Where("id = ?", id).Find(&managerList).RowsAffected
	if affected == 0 {
		con.Error(c, "管理员不存在")
		return
	}
	err = models.DB.Where("id = ?", id).Delete(&models.Manager{}).Error
	if err != nil {
		con.Error(c, "删除数据失败")
		return
	}
	con.Success(c, "删除数据成功")
}

package admin

import (
	"ginDemo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessController struct {
	BaseController
}

// Index 获取权限列表
func (con AccessController) Index(c *gin.Context) {
	var accessList []models.Access
	models.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)
	c.JSON(http.StatusOK, gin.H{
		"accessList": accessList,
	})
}

// Create 添加权限
func (con AccessController) Create(c *gin.Context) {
	moduleName := strings.Trim(c.Query("module_name"), " ")
	actionName := strings.Trim(c.Query("action_name"), " ")
	accessType, err1 := models.StringToInt(c.Query("type"))
	url := c.Query("url")
	moduleId, err2 := models.StringToInt(c.Query("module_id"))
	sort, err3 := models.StringToInt(c.Query("sort"))
	status, err4 := models.StringToInt(c.Query("status"))
	description := strings.Trim(c.Query("description"), " ")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "传入参数错误")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空")
		return
	}
	//添加
	err := models.DB.Create(&models.Access{
		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        accessType,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Status:      status,
		Description: description,
	}).Error
	if err != nil {
		con.Error(c, "权限添加失败")
	}
	con.Success(c, "权限添加成功")

}

// Update 编辑权限
func (con AccessController) Update(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.StringToInt(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入数据错误"+models.IntToString(id))
		return
	}
	//获取表单数据
	moduleName := strings.Trim(c.Query("module_name"), " ")
	actionName := strings.Trim(c.Query("action_name"), " ")
	accessType, err1 := models.StringToInt(c.Query("type"))
	url := c.Query("url")
	moduleId, err2 := models.StringToInt(c.Query("module_id"))
	sort, err3 := models.StringToInt(c.Query("sort"))
	status, err4 := models.StringToInt(c.Query("status"))
	description := strings.Trim(c.Query("description"), " ")
	//判断err
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "传入数据错误"+models.IntToString(id))
		return
	}
	//判断moduleName
	if moduleName == "" {
		con.Error(c, "模块名称不能为空"+models.IntToString(id))
		return
	}

	// 修改
	err = models.DB.Where("id = ?", id).Updates(&models.Access{
		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        accessType,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Status:      status,
		Description: description,
	}).Error
	if err != nil {
		con.Error(c, "修改失败")
		return
	}
	con.Success(c, "修改成功")
}

// Delete 删除权限
func (con AccessController) Delete(c *gin.Context) {
	//获取提交的表单数据
	id, err := models.StringToInt(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误")
		return
	}
	// 删除
	access := models.Access{Id: id}
	models.DB.Find(&access)
	if access.ModuleId == 0 { // 顶级模块
		var accessList []models.Access
		models.DB.Where("module_id = ? ", access.Id).Find(&accessList)
		if len(accessList) > 0 {
			con.Error(c, "当前模块下有子菜单,请先删除子菜单后再来删除这个数据")
			return
		}
	}
	// 顶级模块下面没有子菜单, 可以直接删除
	err = models.DB.Delete(&access).Error
	if err != nil {
		con.Error(c, "删除数据失败")
		return
	}
	con.Success(c, "删除数据成功")
}

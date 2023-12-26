package models

// 角色模型

type Role struct {
	Id          int
	Title       string
	Description string
	Status      int
	AddTime     int
}

// TableName 配置数据库表名称
func (Role) TableName() string {
	return "role"
}

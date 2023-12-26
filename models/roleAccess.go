package models

// Role Access 多对多关系的第三张表

type RoleAccess struct {
	Id       int
	AccessId int
	RoleId   int
}

func (RoleAccess) TableName() string {
	return "role_access"
}

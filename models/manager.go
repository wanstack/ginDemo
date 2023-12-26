package models

//管理员表

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	Role     Role `gorm:"foreignKey:RoleId;references:Id"` //  字段RoleID是本表外键,和role表的id建立外键关系
}

//TableName 配置数据库操作的表名称
func (Manager) TableName() string {
	return "manager"
}

package models

import "fmt"

// CreateTable 创建表, 待优化
func CreateTable() {
	M := DB.Migrator()
	if ok := M.HasTable(&Role{}); ok {
		//M.DropTable(&Role{})
		fmt.Println("role table is exist")
	} else {
		M.CreateTable(&Role{})
	}

	if ok := M.HasTable(&Manager{}); ok {
		//M.DropTable(&Role{})
		fmt.Println("Manager table is exist")
	} else {
		M.CreateTable(&Manager{})
	}

	if ok := M.HasTable(&Access{}); ok {
		//M.DropTable(&Role{})
		fmt.Println("Access table is exist")
	} else {
		M.CreateTable(&Access{})
	}

	if ok := M.HasTable(&RoleAccess{}); ok {
		//M.DropTable(&Role{})
		fmt.Println("RoleAccess table is exist")
	} else {
		M.CreateTable(&RoleAccess{})
	}

}

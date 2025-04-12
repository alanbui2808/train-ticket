package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model        // default fields
	ID         int64  `gorm:"column:id; type:int; non null; primaryKey; autoIncrement; comment:'PK is ID'"`
	RoleName   string `gorm:"column:role_name"`
	RoleNote   string `gorm:"column:role_note; type:text;"` // User-Role: M:N relationship
}

func (r *Role) TableName() string {
	return "go_db_role"
}

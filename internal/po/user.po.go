package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
1. Persistent Object (PO)
Data object whose state is stored to perm storage (like SQL), survives beyond
lifetime of program.

2. Why define Go struct when we already stored data in SQL?
  - Bridge between Go and SQL: Struct can map to SQL objects. Easier to work
    with instead of raw SQL queries.
  - Allow ORMS like GORM to handle CRUD operations automatically. No need to write
    raw SQL queries
*/
type User struct {
	gorm.Model           // default fields
	UUID       uuid.UUID `gorm:"column:uuid; type:varchar(255); non null; index:idx_uuid; unique;"`
	UserName   string    `gorm:"column:user_name"`
	IsActive   bool      `gorm:"column:is_active; type:boolean;"`
	Roles      []Role    `gorm:"many2many:go_user_roles"` // User-Role: M:N relationship
}

func (u *User) TableName() string {
	return "go_db_user"
}

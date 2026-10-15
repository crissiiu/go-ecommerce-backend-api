package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; not null; primarykey; autoIncrement; comment:'Primary key is int';"`
	RoleName string `gorm:"column:role_name;"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}

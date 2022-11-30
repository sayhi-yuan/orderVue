package model

import "time"

type SysUserModel struct {
	ID            int       `json:"id" gorm:"primaryKey;column:id;type:int(11) auto_increment"`
	Name          string    `json:"name" gorm:"column:name;type:varchar(255);not null;comment:姓名"`
	UserName      string    `json:"user_name" gorm:"column:user_name;type:varchar(255);not null;comment:用户名"`
	UserPassword  string    `json:"user_password" gorm:"column:user_password;type:varchar(255);not null;comment:密码"`
	Administrator int       `json:"administrator" gorm:"column:administrator;type:int(11);not null;comment:1管理员,2普通用户"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`
}

// TableName returns the table name of the SysUser model
func (SysUserModel) TableName() string {
	return "sys_user"
}

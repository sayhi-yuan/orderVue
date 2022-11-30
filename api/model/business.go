package model

type SysBusinessModel struct {
	ID              int    `json:"id" gorm:"primaryKey;column:id;type:int(11) auto_increment"`
	BusinessName    string `json:"business_name" gorm:"column:business_name;type:varchar(255);not null;uniqueIndex:un_business_name;comment:商家名称"`
	BusinessAddress string `json:"business_address" gorm:"column:business_address;type:varchar(255);not null;comment:商家地址"`
}

// TableName returns the table name of the SysBusiness model
func (SysBusinessModel) TableName() string {
	return "sys_business"
}

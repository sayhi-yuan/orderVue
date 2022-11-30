package model

type SysColorModel struct {
	ID    int    `json:"id" gorm:"primaryKey;column:id;type:int(11) auto_increment"`
	Color string `json:"color" gorm:"column:color;type:varchar(255);not null"`
}

// TableName returns the table name of the SysColor model
func (SysColorModel) TableName() string {
	return "sys_color"
}

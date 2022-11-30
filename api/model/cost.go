package model

import "time"

type SysNewCostModel struct {
	ID         int       `json:"id" gorm:"primaryKey;column:id;type:int(11) auto_increment"`
	NewCost    int       `json:"new_cost" gorm:"column:new_cost;type:int(11);not null;comment:新增成本"`
	NewReason  string    `json:"new_reason" gorm:"column:new_reason;type:varchar(255);not null;comment:新增成本的理由"`
	ClearCost  int       `json:"clear_cost" gorm:"column:clear_cost;type:int(11);not null;default:0;comment:是否清空成本,0不清空,1清空"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:datetime;not null"`
}

// TableName returns the table name of the SysNewCost model
func (SysNewCostModel) TableName() string {
	return "sys_new_cost"
}

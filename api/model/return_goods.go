package model

import "time"

type SysReturnGoodsModel struct {
	ID              int       `json:"id" gorm:"primaryKey;column:id;type:int(11) auto_increment"`
	BusinessName    string    `json:"business_name" gorm:"column:business_name;type:varchar(255);not null;comment:商家名称"`
	BusinessPrice   int       `json:"business_price" gorm:"column:business_price;type:int(11);not null;comment:商家价格"`
	BusinessNumber  string    `json:"business_number" gorm:"column:business_number;type:varchar(255);not null;comment:商家货号"`
	BusinessColor   string    `json:"business_color" gorm:"primaryKey;column:business_color;type:varchar(255);comment:商家颜色"`
	BusinessAddress string    `json:"business_address" gorm:"column:business_address;type:varchar(255);not null;comment:商家地址"`
	IsView          int       `json:"is_view" gorm:"column:is_view;type:int(11);not null;default:0;comment:是否展示"`
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time;type:datetime;not null;comment:创建时间"`
}

// TableName returns the table name of the SysReturnGoods model
func (SysReturnGoodsModel) TableName() string {
	return "sys_return_goods"
}

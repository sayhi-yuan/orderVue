package model

import "time"

type SysDeliverGoodsModel struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id;type:int(11) auto_increment"`
	CarNumber string    `json:"car_number" gorm:"column:car_number;type:varchar(255);not null;comment:车号"`
	Packet    string    `json:"packet" gorm:"column:packet;type:varchar(255);not null;comment:封包"`
	Packages  string    `json:"packages" gorm:"column:packages;type:varchar(255);not null;comment:包裹"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`
}

// TableName returns the table name of the SysDeliverGoods model
func (SysDeliverGoodsModel) TableName() string {
	return "sys_deliver_goods"
}

package model

import "time"

type SysGoodsModel struct {
	ID                      uint64    `json:"id" gorm:"primaryKey;column:id;type:bigint(20) unsigned auto_increment"`
	SupplierItemNumber      string    `json:"supplier_item_number" gorm:"column:supplier_item_number;type:varchar(100);not null;uniqueIndex:un_supplier_item_number;comment:供方货号"`
	MerchantName            string    `json:"merchant_name" gorm:"column:merchant_name;type:varchar(50);not null;comment:商家名称"`
	MerchantItemNumber      string    `json:"merchant_item_number" gorm:"column:merchant_item_number;type:varchar(100);not null;comment:商家货号"`
	MerchantItemNumberColor string    `json:"merchant_item_number_color" gorm:"column:merchant_item_number_color;type:varchar(50);not null;comment:商家货号颜色"`
	MerchantPrice           int       `json:"merchant_price" gorm:"column:merchant_price;type:int(11);not null;default:0;comment:商家价格"`
	SalePrice               float32   `json:"sale_price" gorm:"column:sale_price;type:float(10,2);not null;default:0.01;comment:售价"`
	MerchantSize            string    `json:"merchant_size" gorm:"column:merchant_size;type:varchar(255);not null;comment:商品尺码"`
	CreatedAt               time.Time `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt               time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	MerchantAddress         string    `json:"merchant_address" gorm:"column:merchant_address;type:varchar(255);default:*****"`
	ImgAdd                  string    `json:"img_add" gorm:"column:img_add;type:text;comment:图片地址，json格式"`
	ClearCost               int       `json:"clear_cost" gorm:"column:clear_cost;type:int(11);not null;default:0"`
}

// TableName returns the table name of the SysGoods model
func (SysGoodsModel) TableName() string {
	return "sys_goods"
}

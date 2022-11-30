package model

import "time"

type SysOrderModel struct {
	ID                      uint64    `json:"id" gorm:"primaryKey;column:id;type:bigint(20) unsigned auto_increment"`
	SupplierItemNumber      string    `json:"supplier_item_number" gorm:"column:supplier_item_number;type:varchar(100);not null;comment:供方货号"`
	MerchantName            string    `json:"merchant_name" gorm:"column:merchant_name;type:varchar(50);not null;comment:商家名称"`
	MerchantItemNumber      string    `json:"merchant_item_number" gorm:"column:merchant_item_number;type:varchar(100);not null;comment:商家货号"`
	MerchantItemNumberColor string    `json:"merchant_item_number_color" gorm:"column:merchant_item_number_color;type:varchar(50);not null;comment:商家货号颜色"`
	MerchantSize            int       `json:"merchant_size" gorm:"column:merchant_size;type:int(11);not null;comment:尺码"`
	MerchantNumber          int       `json:"merchant_number" gorm:"column:merchant_number;type:int(11);not null;comment:数量"`
	MerchantPrice           int       `json:"merchant_price" gorm:"column:merchant_price;type:int(11);not null;default:0;comment:商家价格"`
	SaleOrder               string    `json:"sale_order" gorm:"column:sale_order;type:varchar(255);not null;default:1;comment:前台订单"`
	MerchantAddress         string    `json:"merchant_address" gorm:"column:merchant_address;type:varchar(255);not null;comment:商家地址"`
	MerchantQrCode          string    `json:"merchant_qr_code" gorm:"column:merchant_qr_code;type:varchar(255);comment:二维码地址"`
	IsView                  int       `json:"is_view" gorm:"column:is_view;type:int(2);not null;default:0;comment:是否展示,0展示，1不展示"`
	IsExport                int       `json:"is_export" gorm:"column:is_export;type:int(11);not null;default:1;comment:是否导出0导出，1未导出"`
	ClearCost               int       `json:"clear_cost" gorm:"column:clear_cost;type:int(11);not null;default:0;comment:是否清空成本,0不清空,1清空"`
	CreateTime              time.Time `json:"create_time" gorm:"column:create_time;type:datetime;comment:创建时间"`
}

// TableName returns the table name of the SysOrder model
func (SysOrderModel) TableName() string {
	return "sys_order"
}

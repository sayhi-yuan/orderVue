package req

type GoodsAddReq struct {
	SupplierItemNumber      string  `json:"supplier_item_number"`
	MerchantName            string  `json:"merchant_name"`
	MerchantItemNumber      string  `json:"merchant_item_number" `
	MerchantItemNumberColor string  `json:"merchant_item_number_color"`
	MerchantPrice           int     `json:"merchant_price"`
	SalePrice               float32 `json:"sale_price"`
	MerchantSize            string  `json:"merchant_size"`
	MerchantAddress         string  `json:"merchant_address"`
	ImgAdd                  string  `json:"img_add"`
	ClearCost               int     `json:"clear_cost"`
}

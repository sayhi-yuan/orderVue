package service

import (
	"order/api/errs"
	"order/api/model"
	"order/api/pkg/gormx"
	"order/api/pkg/resp"
	"order/api/req"
)

func (s *Service) GoodsList(name string, page, size int) resp.Response {
	var business []model.SysGoodsModel
	var total int64

	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	offset := (page - 1) * size
	if err := gormx.Db.Where("merchant_name like", "%"+name+"%").
		Count(&total).
		Offset(offset).
		Limit(size).
		Find(&business).Error; err != nil {
		return resp.Error(errs.GetError)
	}

	return resp.SuccessList(total, page, size, business)
}

func (s *Service) GoodsAdd(params req.GoodsAddReq) resp.Response {
	var goods model.SysGoodsModel

	if err := gormx.Db.Where("merchant_name", params.MerchantName).Find(&goods).Error; err != nil {
		return resp.Error(errs.GetError)
	}
	if goods.ID > 0 {
		return resp.Error(errs.ExistError)
	}

	if err := gormx.Db.Create(&goods).Error; err != nil {
		return resp.Error(errs.AddError)
	}

	return resp.Success(goods)
}

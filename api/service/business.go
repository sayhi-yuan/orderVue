package service

import (
	"order/api/errs"
	"order/api/model"
	"order/api/pkg/gormx"
	"order/api/pkg/resp"
)

func (s *Service) BusinessAdd(name, address string) resp.Response {
	var business model.SysBusinessModel

	if err := gormx.Db.Where("business_name", name).Find(&business).Error; err != nil {
		return resp.Error(errs.GetError)
	}

	if business.ID > 0 {
		return resp.Error(errs.ExistError)
	}

	business.BusinessName = name
	business.BusinessAddress = address
	if err := gormx.Db.Create(&business).Error; err != nil {
		return resp.Error(errs.AddError)
	}

	return resp.Success(business)
}

func (s *Service) BusinessList(name string, page, size int) resp.Response {
	var business []model.SysBusinessModel
	var total int64

	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	offset := (page - 1) * size
	if err := gormx.Db.Where("business_name like", "%"+name+"%").
		Count(&total).
		Offset(offset).
		Limit(size).
		Find(&business).Error; err != nil {
		return resp.Error(errs.GetError)
	}

	return resp.SuccessList(total, page, size, business)
}

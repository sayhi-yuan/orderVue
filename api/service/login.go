package service

import (
	"order/api/errs"
	"order/api/model"
	"order/api/pkg/gormx"
	"order/api/pkg/helper"
	"order/api/pkg/resp"
)

func (s *Service) Login(name, pass, database string) resp.Response {
	err := gormx.Connect(gormx.Config{
		Database:           database,
		MaxIdleConnections: 10,
		MaxOpenConnections: 20,
		MaxLifeTime:        10,
	})
	if err != nil {
		return resp.Error(errs.MysqlError)
	}

	var user model.SysUserModel
	if err = gormx.Db.Where("user_name", name).Find(&user).Error; err != nil {
		return resp.Error(errs.GetError)
	}

	if user.ID == 0 {
		return resp.Error(errs.NotExistError)
	}

	if user.UserPassword != helper.Md5V(pass) {
		return resp.Error(errs.UserPassError)
	}

	result := struct {
		Id       int    `json:"id"`
		User     string `json:"user"`
		UserName string `json:"user_name"`
	}{
		Id:       user.ID,
		User:     user.Name,
		UserName: user.UserName,
	}

	return resp.Success(result)
}

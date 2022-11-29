package service

import (
	"log"
	"order/api/pkg/gormx"
	"order/api/pkg/resp"
)

func (s *Service) Login(name, pass, database string) any {
	log.Println("js调用过来了- 参数：", name, pass, database)

	err := gormx.Connect(gormx.Config{
		Database:           database,
		MaxIdleConnections: 10,
		MaxOpenConnections: 20,
		MaxLifeTime:        10,
	})
	if err != nil {
		return resp.Notice(101, "连接数据库异常")
	}

	return resp.Success(map[string]interface{}{"id": 1, "name": "123"})
}

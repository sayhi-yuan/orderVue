package errs

import "order/api/pkg/errorx"

var (
	MysqlError    = errorx.NewErrorMsg(101, "数据库错误")
	GetError      = errorx.NewErrorMsg(102, "获取数据异常")
	ExistError    = errorx.NewErrorMsg(103, "数据已存在")
	NotExistError = errorx.NewErrorMsg(104, "记录未找到")
	AddError      = errorx.NewErrorMsg(105, "添加记录错误")
	UpdateError   = errorx.NewErrorMsg(106, "更新记录错误")

	UserPassError = errorx.NewErrorMsg(203, "用户或密码错误")
)

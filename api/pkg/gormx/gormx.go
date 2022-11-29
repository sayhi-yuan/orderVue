package gormx

import (
	"database/sql"
	"fmt"
	"log"
	"order/api/pkg/logging"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config 配置
type Config struct {
	// dsn
	Database string
	// 最大空闲连接数
	MaxIdleConnections int
	// 最大打开连接数
	MaxOpenConnections int
	// 最长活跃时间
	MaxLifeTime int
}

var Db *gorm.DB

func Connect(c Config) error {
	dnDsn := fmt.Sprintf("root:123456@tcp(localhost:3306)/%s?multiStatements=true&parseTime=true", c.Database)
	sqlDB, err := sql.Open("mysql", dnDsn)
	if err != nil {
		return err
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(c.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(c.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifeTime))

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 multiLogger(),
	})
	if err != nil {
		return err
	}
	Db = db
	return nil
}

func multiLogger() logger.Interface {
	return logger.New(log.New(logging.LogFile, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             time.Second,   // 慢 SQL 阈值
		LogLevel:                  logger.Silent, // 日志级别
		IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,         // 禁用彩色打印
	})
}

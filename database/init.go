package database

import (
	"fmt"
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

// Init 初始化 MySQL 链接
func Init() {
	utils.Log().Info("初始化数据库连接")

	var (
		db  *gorm.DB
		err error
	)

	if gin.Mode() == gin.TestMode {
		// 测试模式下，使用内存数据库
		db, err = gorm.Open(sqlite.Open(":memory:"))
	} else {
		logLevel := logger.Silent

		// Debug模式下，输出所有 SQL 日志
		if conf.SystemConfig.Debug {
			logLevel = logger.Info
		}

		newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		})

		gormConfig := &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   conf.DatabaseConfig.TablePrefix,
				SingularTable: true,
			},
			Logger: newLogger,
		}

		switch conf.DatabaseConfig.Type {
		case "UNSET", "sqlite", "sqlite3":
			// 未指定数据库或者明确指定为 sqlite 时，使用 SQLite3 数据库
			db, err = gorm.Open(sqlite.Open(utils.RelativePath(conf.DatabaseConfig.File)), gormConfig)
		case "postgres":
			db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.User,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Database,
				conf.DatabaseConfig.Port)), gormConfig)
		case "mysql", "mssql":
			db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
				conf.DatabaseConfig.User,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.Port,
				conf.DatabaseConfig.Database,
				conf.DatabaseConfig.Charset)), gormConfig)
		default:
			utils.Log().Panic("不支持数据库类型: %s", conf.DatabaseConfig.Type)
		}
	}

	//db.SetLogger(util.Log())
	if err != nil {
		utils.Log().Panic("连接数据库不成功, %s", err)
	}

	DB = db

	//执行迁移
	migration()
}

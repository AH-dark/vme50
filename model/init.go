package model

import (
	"fmt"
	"github.com/AH-dark/random-donate/pkg/conf"
	"github.com/AH-dark/random-donate/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
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

		var dialector gorm.Dialector

		switch conf.DatabaseConfig.Type {
		case "UNSET", "sqlite", "sqlite3":
			// 未指定数据库或者明确指定为 sqlite 时，使用 SQLite3 数据库
			dialector = sqlite.Open(utils.RelativePath(conf.DatabaseConfig.File))
		case "postgres":
			dialector = postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.User,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Database,
				conf.DatabaseConfig.Port))
		case "mysql":
			dialector = mysql.New(mysql.Config{
				DSN: fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
					conf.DatabaseConfig.User,
					conf.DatabaseConfig.Password,
					conf.DatabaseConfig.Host,
					conf.DatabaseConfig.Port,
					conf.DatabaseConfig.Database,
					conf.DatabaseConfig.Charset),
				DefaultStringSize:         256,   // default size for string fields
				DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
				DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
				DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
				SkipInitializeWithVersion: false, // autoconfigure based on currently MySQL version
			})
		case "mssql":
			dialector = sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
				conf.DatabaseConfig.User,
				conf.DatabaseConfig.Password,
				conf.DatabaseConfig.Host,
				conf.DatabaseConfig.Port,
				conf.DatabaseConfig.Database,
			))
		default:
			utils.Log().Panic("不支持数据库类型: %s", conf.DatabaseConfig.Type)
			return
		}

		db, err = gorm.Open(dialector, gormConfig)
	}

	// db.SetLogger(util.Log())
	if err != nil {
		utils.Log().Panic("连接数据库不成功, %s", err)
	}

	DB = db

	//执行迁移
	migration()
}

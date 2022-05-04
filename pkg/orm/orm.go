package orm

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewOrm 链接数据库，生成数据库实例
func NewOrm(c *Config) (db *gorm.DB) {
	gconfig := gormConfig(c)

	var err error
	switch c.Type {
	case "sqlite3":
		{
			if !(strings.HasSuffix(c.DBFile, ".db") && len(c.DBFile) > 3) {
				log.Panicf("db name error.")
			}
			db, err = gorm.Open(sqlite.Open(c.DBFile), gconfig)
			if err != nil {
				log.Panicf("failed to connect database:%s", err.Error())
			}
		}
	case "mysql":
		{
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				c.Username, c.Password, c.Host, c.Port, c.Dbname)
			db, err = gorm.Open(mysql.Open(dsn), gconfig)
			if err != nil {
				log.Panicf("failed to connect database:%s", err.Error())
			}
		}
	case "postgres":
		{
			dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
				c.Host, c.Username, c.Password, c.Dbname, c.Port)
			db, err = gorm.Open(postgres.Open(dsn), gconfig)
			if err != nil {
				log.Panicf("failed to connect database:%s", err.Error())
			}
		}
	default:
		log.Panicf("not supported database type: %s", c.Type)
	}

	return db
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(c *Config) *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true} // 禁止外键约束, 生产环境不建议使用外键约束
	// 打印所有SQL
	if c.ShowLog {
		config.Logger = logger.Default.LogMode(logger.Info)
	} else {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	return config

}

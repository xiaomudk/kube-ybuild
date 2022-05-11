package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/xiaomudk/kube-ybuild/internal/config"
	"github.com/xiaomudk/kube-ybuild/pkg/orm"
)

// DB 数据库全局变量
var DB *gorm.DB

// Init 初始化数据库
func Init() *gorm.DB {

	DB = orm.NewOrm(config.Conf.Database)
	return DB
}

// GetDB 返回默认的数据库
func GetDB() *gorm.DB {
	return DB
}

// MigrateDatabase run auto migration for given models, will only add missing fields,
// won't delete/change current data.
// nolint:unused // may be reused in the feature, or just show a migrate usage.
func MigrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(UserModel{}); err != nil {
		return errors.Wrap(err, "migrate user model failed")
	}
	return nil
}

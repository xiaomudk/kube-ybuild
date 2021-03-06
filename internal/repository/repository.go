package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/xiaomudk/kube-ybuild/internal/model"
)

var (
	// ErrNotFound data is not exist
	ErrNotFound = gorm.ErrRecordNotFound
)

var _ Repository = (*repository)(nil)

// Repository 定义用户仓库接口
type Repository interface {
	// BaseUser
	CreateUser(ctx context.Context, user *model.UserModel) (id uint64, err error)
	UpdateUser(ctx context.Context, id uint64, userMap map[string]interface{}) error
	DeleteUser(ctx context.Context, id uint64) error
	GetUser(ctx context.Context, id uint64) (*model.UserModel, error)
	GetUsersByIds(ctx context.Context, ids []uint64) ([]*model.UserModel, error)
	GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error)
	UserIsExist(user *model.UserModel) (bool, error)

	Close()
}

// repository mysql struct
type repository struct {
	orm *gorm.DB
}

// New a repository and return
func New(db *gorm.DB) Repository {
	return &repository{
		orm: db,
	}
}

// Close release mysql connection
func (d *repository) Close() {

}

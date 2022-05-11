package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/xiaomudk/kube-ybuild/internal/model"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

// CreateUser 创建用户
func (d *repository) CreateUser(ctx context.Context, user *model.UserModel) (id uint64, err error) {
	err = d.orm.Create(&user).Error
	if err != nil {
		//prom.BusinessErrCount.Incr("mysql: CreateUser")
		return 0, errors.Wrap(err, "[repo.user] create user err")
	}

	return user.ID, nil
}

// UpdateUser 更新用户信息
func (d *repository) UpdateUser(ctx context.Context, id uint64, userMap map[string]interface{}) error {
	user, err := d.GetUser(ctx, id)
	if err != nil {
		//prom.BusinessErrCount.Incr("mysql: getOneUser")
		return errors.Wrap(err, "[repo.user] update user data err")
	}

	err = d.orm.Model(user).Updates(userMap).Error
	if err != nil {
		//prom.BusinessErrCount.Incr("orm: UpdateUser")
	}
	return err
}

// DeleteUser delete user
func (d *repository) DeleteUser(ctx context.Context, id uint64) error {
	user, err := d.GetUser(ctx, id)
	if err != nil {
		//prom.BusinessErrCount.Incr("orm: getOneUser")
		return errors.Wrap(err, "[repo.user] delete user data err")
	}

	err = d.orm.Delete(user).Error

	if err != nil {
		//prom.BusinessErrCount.Incr("orm: DeleteUser")
	}
	return err
}

// GetUser 获取用户
func (d *repository) GetUser(ctx context.Context, uid uint64) (user *model.UserModel, err error) {

	// 从数据库中获取
	err = d.orm.WithContext(ctx).First(&user, uid).Error
	if errors.Is(err, ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, errors.Wrapf(err, "[repo.user] query db err")
	}

	return user, nil
}

// GetUsersByIds 批量获取用户
func (d *repository) GetUsersByIds(ctx context.Context, userIDs []uint64) ([]*model.UserModel, error) {
	users := make([]*model.UserModel, 0)

	// 查询未命中
	for _, userID := range userIDs {
		userModel, err := d.GetUser(ctx, userID)
		if err != nil {
			logs.Warnf("[repo.user_base] get user model err: %v", err)
			continue
		}
		users = append(users, userModel)
	}
	return users, nil
}

// GetUserByPhone 根据手机号获取用户
func (d *repository) GetUserByPhone(ctx context.Context, phone int64) (*model.UserModel, error) {
	user := model.UserModel{}
	err := d.orm.Where("phone = ?", phone).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user_base] get user err by phone")
	}

	return &user, nil
}

// GetUserByUsername 根据用户名获取用户
func (d *repository) GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	user := model.UserModel{}
	err := d.orm.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[repo.user] get user err by username")
	}

	return &user, nil
}

// UserIsExist 判断用户是否存在, 用户名和邮箱要保持唯一
func (d *repository) UserIsExist(user *model.UserModel) (bool, error) {
	err := d.orm.Where("username = ? or email = ?", user.Username, user.Email).First(&model.UserModel{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

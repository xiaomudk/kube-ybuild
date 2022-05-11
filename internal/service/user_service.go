package service

import (
	"context"
	"time"

	"github.com/xiaomudk/kube-ybuild/internal/repository"
	"github.com/xiaomudk/kube-ybuild/pkg/utils"

	"github.com/pkg/errors"

	"github.com/xiaomudk/kube-ybuild/internal/model"
	"github.com/xiaomudk/kube-ybuild/pkg/auth"
)

// UserService define interface func
type UserService interface {
	Register(ctx context.Context, username, email, password string) error
	Login(ctx context.Context, username, password string) (tokenStr string, err error)
	GetUserByID(ctx context.Context, id uint64) (*model.UserModel, error)
	GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error)
	UpdateUser(ctx context.Context, id uint64, userMap map[string]interface{}) error
	DeleteUser(ctx context.Context, id uint64) error
}

type userService struct {
	repo repository.Repository
}

var _ UserService = (*userService)(nil)

func newUsers(svc *service) *userService {
	return &userService{repo: svc.repo}
}

// Register 注册用户
func (s *userService) Register(ctx context.Context, username, email, password string) error {
	pwd, err := auth.HashAndSalt(password)
	if err != nil {
		return errors.Wrapf(err, "encrypt password err")
	}

	u := model.UserModel{
		Username:  username,
		Password:  pwd,
		Email:     email,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	isExist, err := s.repo.UserIsExist(&u)
	if err != nil {
		return errors.Wrapf(err, "create user")
	}
	if isExist {
		return errors.New("用户已存在")
	}
	_, err = s.repo.CreateUser(ctx, &u)
	if err != nil {
		return errors.Wrapf(err, "create user")
	}
	return nil
}

// Login 登录
func (s *userService) Login(ctx context.Context, email, password string) (tokenStr string, err error) {
	u, err := s.GetUserByUsername(ctx, email)
	if err != nil {
		return "", errors.Wrapf(err, "get user info err by email")
	}

	// ComparePasswords the login password with the user password.
	if !auth.ComparePasswords(u.Password, password) {
		return "", errors.New("invalid password")
	}

	// 签发签名 Sign the json web token.
	payload := map[string]interface{}{"user_id": u.ID, "username": u.Username}
	tokenStr, err = utils.Sign(payload, "secret", 86400)
	if err != nil {
		return "", errors.Wrapf(err, "gen token sign err")
	}

	return tokenStr, nil
}

// UpdateUser update user info
func (s *userService) UpdateUser(ctx context.Context, id uint64, userMap map[string]interface{}) error {
	err := s.repo.UpdateUser(ctx, id, userMap)

	if err != nil {
		return err
	}

	return nil
}

// DeleteUser delete user
func (s *userService) DeleteUser(ctx context.Context, id uint64) error {
	err := s.repo.DeleteUser(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

// GetUserByID 获取单条用户信息
func (s *userService) GetUserByID(ctx context.Context, id uint64) (*model.UserModel, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	userModel, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return userModel, errors.Wrapf(err, "get user info err from db by username: %s", username)
	}

	return userModel, nil
}

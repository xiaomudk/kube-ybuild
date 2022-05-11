package ecode

import "github.com/xiaomudk/kube-ybuild/pkg/errcode"

//nolint: golint
var (
	// user errors
	ErrUserNotFound          = errcode.NewError(20101, "The user was not found.")
	ErrPasswordIncorrect     = errcode.NewError(20102, "账号或密码错误")
	ErrTwicePasswordNotMatch = errcode.NewError(20103, "两次密码输入不一致")
	ErrRegisterFailed        = errcode.NewError(20104, "注册失败")
)

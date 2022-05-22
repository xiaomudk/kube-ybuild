package ecode

import "github.com/xiaomudk/kube-ybuild/pkg/errcode"

//nolint: golint
var (
	// user errors
	ErrTemplateNotFound = errcode.NewError(30101, "The template was not found.")
	ErrTemplateCreate   = errcode.NewError(30102, "create template error")
)

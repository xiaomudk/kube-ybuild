package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/xiaomudk/kube-ybuild/internal/model"
	"github.com/xiaomudk/kube-ybuild/internal/repository"
	"time"
)

type TemplateService interface {
	CreateTemplate(ctx context.Context, tplName, tplContent string) error
	UpdateTemplate(ctx context.Context, id uint64, templateMap map[string]interface{}) error
	GetTemplateById(ctx context.Context, id uint64) (tpl *model.TemplateModel, err error)
	DeleteTemplate(ctx context.Context, id uint64) error
	ListTemplate(ctx context.Context) ([]model.TemplateModel, error)
	SearchTemplateByName(ctx context.Context, tplName string) ([]model.TemplateModel, error)
}

type templateService struct {
	repo repository.Repository
}

var _ TemplateService = (*templateService)(nil)

func newTemplates(svc *service) *templateService {
	return &templateService{repo: svc.repo}
}

func (t *templateService) CreateTemplate(ctx context.Context, tplName, tplContent string) error {
	tpl := model.TemplateModel{
		TemplateName:    tplName,
		TemplateContent: tplContent,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}
	_, err := t.repo.CreateTemplate(ctx, &tpl)
	if err != nil {
		return errors.Wrapf(err, "create template")
	}
	return nil
}

func (t *templateService) UpdateTemplate(ctx context.Context, id uint64, templateMap map[string]interface{}) error {
	err := t.repo.UpdateTemplate(ctx, id, templateMap)
	if err != nil {
		return err
	}
	return nil
}

func (t *templateService) GetTemplateById(ctx context.Context, id uint64) (tpl *model.TemplateModel, err error) {
	tpl, err = t.repo.GetTemplate(ctx, id)
	if err != nil {
		return nil, err
	}
	return tpl, nil
}

func (t *templateService) DeleteTemplate(ctx context.Context, id uint64) error {
	err := t.repo.DeleteTemplate(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

func (t *templateService) ListTemplate(ctx context.Context) ([]model.TemplateModel, error) {
	tpl, err := t.repo.ListTemplate(ctx)
	if err != nil {
		return nil, err
	}
	return tpl, nil
}

func (t *templateService) SearchTemplateByName(ctx context.Context, tplName string) ([]model.TemplateModel, error) {
	tpl, err := t.repo.SearchTemplateByName(ctx, tplName)
	if err != nil {
		return nil, err
	}
	return tpl, nil
}

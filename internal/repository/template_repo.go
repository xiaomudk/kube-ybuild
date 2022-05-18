package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/xiaomudk/kube-ybuild/internal/model"
)

func (d *repository) CreateTemplate(ctx context.Context, tpl *model.TemplateModel) (id uint64, err error) {
	err = d.orm.Create(&tpl).Error
	if err != nil {
		return 0, errors.Wrap(err, "[repo.template] create template err")
	}

	return tpl.Id, nil
}

func (d *repository) UpdateTemplate(ctx context.Context, id uint64, templateMap map[string]interface{}) error {
	tpl, err := d.GetTemplate(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[repo.template] get template id err")
	}
	err = d.orm.Model(tpl).Updates(templateMap).Error
	if err != nil {
		return errors.Wrap(err, "[repo.template] update template data err")
	}
	return err
}

func (d *repository) GetTemplate(ctx context.Context, id uint64) (tpl *model.TemplateModel, err error) {
	// 从数据库取template
	err = d.orm.WithContext(ctx).First(&tpl, id).Error
	if errors.Is(err, ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, errors.Wrapf(err, "[repo.template] query db err")
	}
	return tpl, nil
}

func (d *repository) GetTemplateByTplName(ctx context.Context, tplName string) ([]model.TemplateModel, error) {
	var tplList []model.TemplateModel
	err := d.orm.Where("tpl_name LIKE ?", "%"+tplName+"%").Find(&tplList).Error
	if err != nil {
		return tplList, errors.Wrap(err, "[repo.template] get template err by tplName")
	}
	return tplList, nil
}

func (d *repository) DeleteTemplate(ctx context.Context, id uint64) error {
	tpl, err := d.GetTemplate(ctx, id)
	if err != nil {
		//prom.BusinessErrCount.Incr("orm: getOneUser")
		return errors.Wrap(err, "[repo.template] delete template data err")
	}

	err = d.orm.Delete(tpl).Error

	if err != nil {
		//prom.BusinessErrCount.Incr("orm: DeleteUser")
	}
	return err
}

func (d *repository) ListTemplate(ctx context.Context) ([]model.TemplateModel, error) {
	var tplList []model.TemplateModel
	err := d.orm.Find(&tplList).Error
	if err != nil {
		return tplList, errors.Wrap(err, "[repo.template] list template err")
	}
	return tplList, nil
}

func (d *repository) SearchTemplateByName(ctx context.Context, tplName string) ([]model.TemplateModel, error) {
	var tplList []model.TemplateModel
	err := d.orm.Where("template_name LIKE ?", "%"+tplName+"%").Find(&tplList).Error
	if err != nil {
		return tplList, errors.Wrap(err, "[repo.template] search template err")
	}
	return tplList, nil
}

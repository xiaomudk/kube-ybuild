package model

import "time"

type TemplateModel struct {
	Id              uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	TemplateName    string    `gorm:"column:template_name;type:varchar(255);unique;not null;comment:模板名称" binding:"required"  json:"templateName"`
	TemplateContent string    `gorm:"column:template_content;type:longtext;comment:模板内容" json:"templateContent"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (t *TemplateModel) TableName() string {
	return "template"
}

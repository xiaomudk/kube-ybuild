package utils

import (
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

func Paginate(stmt *gorm.DB, req interface{}, res interface{}) *paginate.Page {
	pg := paginate.New()
	page := pg.Response(stmt, req, res)
	return &page
}

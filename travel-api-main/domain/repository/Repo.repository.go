package repository

import (
	"gorm.io/gorm"
)

func Paginate(page int64, size int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		paginate := Pagination(page, size)
		return db.Offset(int(paginate.Offset)).Limit(int(size))
	}
}

package paginate

import (
	"github.com/jassue/gin-wire/app/pkg/request"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

func Paginate(req *request.PageDto) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := req.Page
		if page == 0 {
			page = 1
		}
		pageSize := req.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func MPaginate(req *request.PageDto) *options.FindOptions {

	page := req.Page
	if page == 0 {
		page = 1
	}

	pageSize := req.PageSize

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	findOptions := options.Find()

	//fmt.Println("page", page, "size", pageSize)

	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))
	//findOptions := options.FindOptions{Limit: &limit, Skip: &skip}

	return findOptions
}

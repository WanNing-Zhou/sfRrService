package data

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/model"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/service"
	"github.com/jassue/gin-wire/util/paginate"
	"github.com/jassue/gin-wire/util/toModule"
	"go.uber.org/zap"
)

type compRepo struct {
	data *Data
	log  *zap.Logger
}

func NewCompRepo(data *Data, log *zap.Logger) service.CompRepo {
	return &compRepo{
		data: data,
		log:  log,
	}
}

// 根据id寻找comp

func (r *compRepo) FindByID(ctx context.Context, id uint64) (*domain.Comp, error) {
	var comp model.Comp
	if err := r.data.db.First(&comp, id).Error; err != nil {
		return nil, err
	}
	return comp.ToDomain(), nil
}

// 根据创建者id寻找comp

func (r *compRepo) FindByCreateId(ctx context.Context, creatId uint64) (*domain.Comp, error) {
	var comp model.Comp
	if err := r.data.db.First(&comp, creatId).First(&comp).Error; err != nil {
		return nil, err
	}
	return comp.ToDomain(), nil
}

// Create 床架comp
func (r *compRepo) Create(ctx context.Context, c *domain.Comp) (*domain.Comp, error) {
	var comp model.Comp

	id, err := r.data.sf.NextID()
	if err != nil {
		return nil, err
	}
	toModule.CompDoMainToModule(&comp, c)
	comp.ID = id
	if err = r.data.DB(ctx).Create(&comp).Error; err != nil {
		return nil, err
	}

	return comp.ToDomain(), nil
}

// FindCompsByQuery 查询Comps
func (r *compRepo) FindCompsByQuery(ctx context.Context, param *request.CompList) ([]domain.Comp, int64, error) {
	var comp model.Comp

	query := r.data.db.Model(&comp)

	// 对名字进行模糊查询
	if param.Name != "" {
		query = query.Where("title LIKE ?", "%"+param.Name+"%")
	}
	if param.CreateId != 0 {
		query = query.Where("creat_id = ?", param.CreateId)
	}
	if param.ID != 0 {
		query = query.Where("id = ?", param.ID)
	}

	// 排序
	query = query.Order("updated_at desc")

	pageReq := &request.PageDto{
		Page:     param.Page,
		PageSize: param.PageSize,
	}

	var compList []model.Comp
	var total int64
	// 获取获取过滤后数据的总条数
	query.Count(&total)

	err := query.Scopes(paginate.Paginate(pageReq)).Find(&compList).Error
	if err != nil {
		return nil, 0, err
	}

	var comps []domain.Comp

	for _, v := range compList {
		comps = append(comps, *v.ToDomain())
	}

	return comps, total, nil

}

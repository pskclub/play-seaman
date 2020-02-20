package products

import (
	"bitbucket.org/3dsinteractive/pam4/config"
	"bitbucket.org/3dsinteractive/seaman"
	"golang-guideline/models"
)

type IProductService interface {
	Find(id string, options *FindProductServiceOptions) (interface{}, seaman.IError)
	GetPage(pageOptions *models.PageOptions, options *GetPageProductServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetAfter(pageOptions *models.AfterOptions, options *GetAfterProductServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetBefore(pageOptions *models.BeforeOptions, options *GetBeforeProductServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetAll(options *GetAllProductServiceOptions) ([]interface{}, seaman.IError)
	Delete(id string) seaman.IError
	Update(id string, data interface{}) (interface{}, seaman.IError)
	Create(data interface{}) (interface{}, seaman.IError)
}

func NewProductService(options *ProductServiceOptions) IProductService {
	return &productService{
		ctx:   options.ctx,
		cfg:   options.cfg,
		store: options.store,
	}
}

type productService struct {
	ctx   seaman.IContext
	cfg   config.IConfig
	store IProductStore
}

func (s productService) Find(id string, options *FindProductServiceOptions) (interface{}, seaman.IError) {
	//p.cfg.IndexerConfig().Gzip()
	product, err := s.store.Find(id, &FindProductStoreOptions{
		IsEnabled: options.IsEnabled,
		IsDeleted: options.IsDeleted,
	})

	if err != nil {
		return nil, s.ctx.NewError(err, err, id, options)
	}

	return product, nil
}

func (s productService) GetPage(pageOptions *models.PageOptions, options *GetPageProductServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productService) GetAfter(pageOptions *models.AfterOptions, options *GetAfterProductServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productService) GetBefore(pageOptions *models.BeforeOptions, options *GetBeforeProductServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productService) GetAll(options *GetAllProductServiceOptions) ([]interface{}, seaman.IError) {
	panic("implement me")
}

func (s productService) Delete(id string) seaman.IError {
	panic("implement me")
}

func (s productService) Update(id string, data interface{}) (interface{}, seaman.IError) {
	panic("implement me")
}

func (s productService) Create(data interface{}) (interface{}, seaman.IError) {
	panic("implement me")
}

package products

import (
	"bitbucket.org/3dsinteractive/seaman"
	"golang-guideline/config"
	"golang-guideline/errors"
	"golang-guideline/f"
	"golang-guideline/models"
)

type IProductService interface {
	Find(id string, options *FindProductServiceOptions) (interface{}, seaman.IError)
	GetPage(pageOptions *models.PageOptions, options *GetProductPageServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetAfter(pageOptions *models.AfterOptions, options *GetProductAfterServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetBefore(pageOptions *models.BeforeOptions, options *GetProductBeforeServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetAll(options *GetAllProductServiceOptions) ([]interface{}, seaman.IError)
	Import(csv seaman.CSV, options *ImportProductServiceOptions) ([]interface{}, seaman.IError)
	Delete(id string) (string, seaman.IError)
	Update(id string, data interface{}) (interface{}, seaman.IError)
	Create(data interface{}) (interface{}, seaman.IError)
}

func NewProductService(options *ProductServiceOptions) IProductService {
	return &productService{
		ctx:   options.Context,
		cfg:   options.Config,
		store: options.Store,
	}
}

type productService struct {
	ctx   seaman.IContext
	cfg   config.IConfig
	store IProductStore
}

func (s productService) Import(csv seaman.CSV, options *ImportProductServiceOptions) ([]interface{}, seaman.IError) {
	if err := csv.OpenR(); err != nil {
		return nil, s.ctx.NewError(err, errors.ProductImportFileTypeNotSupport, options)
	}
	panic("implement me")
}

func (s productService) Find(id string, options *FindProductServiceOptions) (interface{}, seaman.IError) {
	//p.cfg.IndexerConfig().Gzip()
	isEnabled := f.GetBoolFalse(options.IsEnabled)
	product, err := s.store.Find(id, &FindProductStoreOptions{
		IsEnabled: &isEnabled,
		IsDeleted: options.IsDeleted,
	})

	if err != nil {
		return nil, s.ctx.NewError(err, err, id, options)
	}

	return product, nil
}

func (s productService) GetPage(pageOptions *models.PageOptions, options *GetProductPageServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	products := make([]interface{}, 0)
	panic("implement me")
	return products, nil, nil
}

func (s productService) GetAfter(pageOptions *models.AfterOptions, options *GetProductAfterServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productService) GetBefore(pageOptions *models.BeforeOptions, options *GetProductBeforeServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productService) GetAll(options *GetAllProductServiceOptions) ([]interface{}, seaman.IError) {
	panic("implement me")
}

func (s productService) Delete(id string) (string, seaman.IError) {
	if id == "" {
		return "", s.ctx.NewError(errors.NotFound, errors.NotFound, id)
	}
	panic("implement me")
}

func (s productService) Update(id string, data interface{}) (interface{}, seaman.IError) {
	panic("implement me")
}

func (s productService) Create(data interface{}) (interface{}, seaman.IError) {
	panic("implement me")
}

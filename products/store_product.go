package products

import (
	"bitbucket.org/3dsinteractive/seaman"
	"golang-guideline/config"
	"golang-guideline/models"
)

type IProductStore interface {
	Find(id string, options *FindProductStoreOptions) (interface{}, seaman.IError)
	GetPage(pageOptions *models.PageOptions, options *GetProductPageStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetAfter(pageOptions *models.AfterOptions, options *GetProductAfterStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetBefore(pageOptions *models.BeforeOptions, options *GetProductBeforeStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetAll(options *GetAllProductStoreOptions) ([]interface{}, seaman.IError)
	Delete(id string) seaman.IError
	Update(id string, data interface{}) (interface{}, seaman.IError)
	Create(data interface{}) (interface{}, seaman.IError)
}

func NewProductStore(options *ProductStoreOptions) IProductStore {
	return &productStore{
		ctx: options.Context,
		cfg: options.Config,
	}
}

type productStore struct {
	ctx seaman.IContext
	cfg config.IConfig
}

func (s productStore) Find(id string, options *FindProductStoreOptions) (interface{}, seaman.IError) {
	panic("implement me")
}

func (s productStore) GetPage(pageOptions *models.PageOptions, options *GetProductPageStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productStore) GetAfter(pageOptions *models.AfterOptions, options *GetProductAfterStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productStore) GetBefore(pageOptions *models.BeforeOptions, options *GetProductBeforeStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productStore) GetAll(options *GetAllProductStoreOptions) ([]interface{}, seaman.IError) {
	panic("implement me")
}

func (s productStore) Delete(id string) seaman.IError {
	panic("implement me")
}

func (s productStore) Update(id string, data interface{}) (interface{}, seaman.IError) {
	panic("implement me")
}

func (s productStore) Create(data interface{}) (interface{}, seaman.IError) {
	panic("implement me")
}

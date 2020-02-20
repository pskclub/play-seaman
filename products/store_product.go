package products

import (
	"bitbucket.org/3dsinteractive/pam4/config"
	"bitbucket.org/3dsinteractive/seaman"
	"golang-guideline/models"
)

type IProductStore interface {
	Find(id string, options *FindProductStoreOptions) (interface{}, seaman.IError)
	GetPage(pageOptions *models.PageOptions, options *GetPageProductStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetAfter(pageOptions *models.AfterOptions, options *GetAfterProductStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetBefore(pageOptions *models.BeforeOptions, options *GetBeforeProductStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	GetAll(options *GetAllProductStoreOptions) ([]interface{}, seaman.IError)
	Delete(id string) seaman.IError
	Update(id string, data interface{}) (interface{}, seaman.IError)
	Create(data interface{}) (interface{}, seaman.IError)
}

func NewProductStore(options *ProductStoreOptions) IProductStore {
	return &productStore{
		ctx: options.ctx,
		cfg: options.cfg,
	}
}

type productStore struct {
	ctx seaman.IContext
	cfg config.IConfig
}

func (s productStore) Find(id string, options *FindProductStoreOptions) (interface{}, seaman.IError) {
	panic("implement me")
}

func (s productStore) GetPage(pageOptions *models.PageOptions, options *GetPageProductStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productStore) GetAfter(pageOptions *models.AfterOptions, options *GetAfterProductStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	panic("implement me")
}

func (s productStore) GetBefore(pageOptions *models.BeforeOptions, options *GetBeforeProductStoreOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
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

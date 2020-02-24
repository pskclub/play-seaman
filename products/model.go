package products

import (
	"bitbucket.org/3dsinteractive/seaman"
	"golang-guideline/config"
)

type GetProductPageServiceOptions struct {
	TagIDs []*string
}

type GetProductAfterServiceOptions struct {
	sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetProductBeforeServiceOptions struct {
	sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetAllProductServiceOptions struct {
}

type ImportProductServiceOptions struct {
}

type ProductServiceOptions struct {
	Context seaman.IContext
	Config  config.IConfig
	Store   IProductStore
}

type FindProductServiceOptions struct {
	IsEnabled *bool
	IsDeleted *bool
}

type ProductStoreOptions struct {
	Context seaman.IContext
	Config  config.IConfig
}

type FindProductStoreOptions struct {
	IsEnabled *bool
	IsDeleted *bool
}

type GetProductPageStoreOptions struct {
	tagIDs []*string
}

type GetProductAfterStoreOptions struct {
	sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetProductBeforeStoreOptions struct {
	sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetAllProductStoreOptions struct {
}

type OrderPayload struct {
}

package products

import (
	"bitbucket.org/3dsinteractive/pam4/config"
	"bitbucket.org/3dsinteractive/seaman"
)

type ProductServiceOptions struct {
	ctx   seaman.IContext
	cfg   config.IConfig
	store IProductStore
}

type FindProductServiceOptions struct {
	IsEnabled *bool
	IsDeleted *bool
}

type GetPageProductServiceOptions struct {
	tagIDs []*string
}

type GetAfterProductServiceOptions struct {
	sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetBeforeProductServiceOptions struct {
	sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetAllProductServiceOptions struct {
}

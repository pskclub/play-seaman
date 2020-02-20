package products

import (
	"bitbucket.org/3dsinteractive/pam4/config"
	"bitbucket.org/3dsinteractive/seaman"
)

type ProductStoreOptions struct {
	ctx seaman.IContext
	cfg config.IConfig
}

type FindProductStoreOptions struct {
	IsEnabled *bool
	IsDeleted *bool
}

type GetPageProductStoreOptions struct {
	tagIDs []*string
}

type GetAfterProductStoreOptions struct {
	sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetBeforeProductStoreOptions struct {
	sort      string
	IsEnabled *bool
	IsDeleted *bool
}

type GetAllProductStoreOptions struct {
}

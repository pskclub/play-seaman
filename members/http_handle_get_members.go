package members

import (
	"bitbucket.org/3dsinteractive/pam4/config"
	"bitbucket.org/3dsinteractive/seaman"
	"golang-guideline/core"
	"golang-guideline/products"
)

func (svc *MemberHTTP) setupGetMembers(e *seaman.Engine) {
	filter := &seaman.Filter{
		Request: &seaman.FilterRequest{
			Endpoint: "/api/members",
			Method:   seaman.HTTPMethodGET,
			Params: []*seaman.AttributeParam{
				paramPageQueryString,
			},
			Attributes: []*seaman.Attribute{
				attributeMemberDescription,
			},
			DefaultPageSize: config.NewEComDataConfig().FrontPageSize(),
			Settings:        seaman.NewRESTAPISettingsArray("items"),
			IsMock:          false,
		},
		Process: &seaman.FilterProcess{
			IndexerConfig: svc.config.IndexerConfig(),
			CacheConfig:   svc.config.CacherConfig(),
			FilterArrayHandler: func(ctx seaman.IContext, data seaman.KeyValueList) (interface{}, *seaman.Pagination, error) {
				service := NewMemberService(&MemberServiceOptions{
					Context: ctx,
					Config:  svc.config,
					ProductService: products.NewProductService(&products.ProductServiceOptions{
						Context: ctx,
						Config:  svc.config,
						Store: products.NewProductStore(&products.ProductStoreOptions{
							Context: ctx,
							Config:  svc.config,
						}),
					}),
				})

				return svc.handleGetMembers(ctx, service)
			},
		},
	}

	e.AddFilter(filter)
}

func (svc *MemberHTTP) handleGetMembers(ctx seaman.IContext, service IMemberService) (interface{}, *seaman.Pagination, error) {
	pageOptions := core.GetPageOptions(ctx.TaskParams(), nil)
	productItems, pageMeta, err := service.GetPage(pageOptions, &GetMemberPageServiceOptions{
		TagIDs: []*string{},
	})

	if err != nil {
		return nil, nil, ctx.NewError(err, err)
	}

	return productItems, seaman.NewPaginationFromPageNumber(pageMeta.Total, pageMeta.Page), nil
}

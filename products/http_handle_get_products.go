package products

import (
	"bitbucket.org/3dsinteractive/seaman"
	"golang-guideline/core"
)

func (svc *ProductHTTP) setupGetProducts(e *seaman.Engine) {
	filter := &seaman.Filter{
		Request: &seaman.FilterRequest{
			Endpoint: "/api/products",
			Method:   seaman.HTTPMethodGET,
			Params: []*seaman.AttributeParam{
				paramPageQueryString,
			},
			Attributes: []*seaman.Attribute{
				attributeProductDescription,
			},
			Settings: seaman.NewRESTAPISettingsArray("items"),
			IsMock:   false,
		},
		Process: &seaman.FilterProcess{
			IndexerConfig: svc.config.IndexerConfig(),
			CacheConfig:   svc.config.CacherConfig(),
			FilterArrayHandler: func(ctx seaman.IContext, data seaman.KeyValueList) (interface{}, *seaman.Pagination, error) {
				service := NewProductService(&ProductServiceOptions{
					Context: ctx,
					Config:  svc.config,
					Store: NewProductStore(&ProductStoreOptions{
						Context: ctx,
						Config:  svc.config,
					}),
				})

				return svc.handleGetProducts(ctx, service)
			},
		},
	}

	e.AddFilter(filter)
}

func (svc *ProductHTTP) handleGetProducts(ctx seaman.IContext, service IProductService) (interface{}, *seaman.Pagination, error) {
	pageOptions := core.GetPageOptions(ctx.TaskParams(), nil)
	products, pageMeta, err := service.GetPage(pageOptions, &GetProductPageServiceOptions{
		TagIDs: []*string{},
	})

	if err != nil {
		return nil, nil, ctx.NewError(err, err)
	}

	return products, seaman.NewPaginationFromPageNumber(pageMeta.Total, pageMeta.Page), nil
}

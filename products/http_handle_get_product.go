package products

import (
	"bitbucket.org/3dsinteractive/pam4/bconfig"
	"bitbucket.org/3dsinteractive/seaman"
)

func (svc *ProductHTTP) setupGetProduct(e *seaman.Engine) {
	filter := &seaman.Filter{
		Request: &seaman.FilterRequest{
			Endpoint: "/api/products/:id",
			Method:   seaman.HTTPMethodGET,
			Params: []*seaman.AttributeParam{
				bconfig.ParamTimestampHeader,
				bconfig.ParamAuthorizationHeader,
				bconfig.ParamIDInURL,
			},
			Attributes: []*seaman.Attribute{},
			Settings:   seaman.NewRESTAPISettings(),
			IsMock:     false,
		},
		Process: &seaman.FilterProcess{
			IndexerConfig: svc.config.IndexerConfig(),
			CacheConfig:   svc.config.CacherConfig(),
			FilterItemHandler: func(ctx seaman.IContext, data seaman.KeyValueList) (interface{}, error) {
				service := NewProductService(&ProductServiceOptions{
					Context: ctx,
					Config:  svc.config,
					Store: NewProductStore(&ProductStoreOptions{
						Context: ctx,
						Config:  svc.config,
					}),
				})
				return svc.handleGetProduct(ctx, service)
			},
		},
	}

	e.AddFilter(filter)
}

func (svc *ProductHTTP) handleGetProduct(ctx seaman.IContext, service IProductService) (interface{}, error) {
	id, _ := ctx.TaskParams().GetString(bconfig.ParamIDInURL.Name, "")

	product, err := service.Find(id, &FindProductServiceOptions{})
	if err != nil {
		return nil, ctx.NewError(err, err, id)
	}

	return product, nil
}

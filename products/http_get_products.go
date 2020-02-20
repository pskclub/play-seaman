// Copyright © PushAndMotion 2018 All Rights Reserved.
// PAM Engine & Library is proprietary and confidential.
// Un-authorize using, editing, copying, adapting, distributing, of this file or some part of this file without
// the prior written consent of PushAndMotion, via any medium is strictly prohibited.
// If not expressively specify in the document, the authorisation to use this library will be granted per application.
// Any question regarding this copyright notice please contact contact@pushandmotion.com
// This copyright notice must be included in the header of every distribution of all the source code.

package products

import (
	"bitbucket.org/3dsinteractive/pam4/config"
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
			DefaultPageSize: config.NewEComDataConfig().FrontPageSize(),
			Settings:        seaman.NewRESTAPISettingsArray("items"),
			IsMock:          false,
		},
		Process: &seaman.FilterProcess{
			IndexerConfig: svc.config.IndexerConfig(),
			CacheConfig:   svc.config.CacherConfig(),
			FilterArrayHandler: func(ctx seaman.IContext, data seaman.KeyValueList) (interface{}, *seaman.Pagination, error) {
				service := NewProductService(&ProductServiceOptions{
					ctx: ctx,
					cfg: svc.config,
					store: NewProductStore(&ProductStoreOptions{
						ctx: ctx,
						cfg: svc.config,
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
	products, pageMeta, err := service.GetPage(pageOptions, &GetPageProductServiceOptions{
		tagIDs: []*string{},
	})

	if err != nil {
		return nil, nil, ctx.NewError(err, err)
	}

	return products, seaman.NewPaginationFromPageNumber(pageMeta.Total, pageMeta.Page), nil
}

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
	"bitbucket.org/3dsinteractive/pam4/services"
	"bitbucket.org/3dsinteractive/seaman"
)

// ProductHTTP is the REST endpoint
type ProductHTTP struct {
	services.BaseHTTPService
	config config.IConfig
}

// NewProductHTTP creates Contact REST handlers
func NewProductHTTP(config config.IConfig) *ProductHTTP {
	return &ProductHTTP{
		config: config,
	}
}

// Name return REST Name
func (svc *ProductHTTP) Name() string {
	return "ProductHTTP"
}

// SetupHandlers setups all handler for the User REST
func (svc *ProductHTTP) SetupHandlers(e *seaman.Engine) {
	svc.setupGetProducts(e) // api
	svc.setupGetProduct(e) // api
}

package products

import (
	"bitbucket.org/3dsinteractive/pam4/config"
	"bitbucket.org/3dsinteractive/pam4/services"
	"bitbucket.org/3dsinteractive/seaman"
)

// CreateShipmentInfoMQ is the MQ endpoint
type ProductMQ struct {
	services.BaseMQService
	config config.IConfig
}

// NewCreateShipmentInfoMQ creates CreateShipmentInfoMQ
func NewProductMQ(config config.IConfig) *ProductMQ {
	return &ProductMQ{
		config: config,
	}
}

// Name return MQ Name
func (svc *ProductMQ) Name() string {
	return "ProductMQ"
}

// SetupHandlers setups all handler
func (svc *ProductMQ) SetupHandlers(e *seaman.Engine) {
	svc.setupCreateShipmentInfoMQ(e)
}

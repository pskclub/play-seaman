// Copyright © PushAndMotion 2018 All Rights Reserved.
// PAM Engine & Library is proprietary and confidential.
// Un-authorize using, editing, copying, adapting, distributing, of this file or some part of this file without
// the prior written consent of PushAndMotion, via any medium is strictly prohibited.
// If not expressively specify in the document, the authorisation to use this library will be granted per application.
// Any question regarding this copyright notice please contact contact@pushandmotion.com
// This copyright notice must be included in the header of every distribution of all the source code.

package products

import (
	"bitbucket.org/3dsinteractive/pam4/errors"
	"bitbucket.org/3dsinteractive/seaman"
)

func (svc *ProductMQ) setupCreateShipmentInfoMQ(e *seaman.Engine) {
	cfg := svc.config.ConsumerConfig()
	e.AddMQConsumer(seaman.NewMQConsumer(
		cfg.Endpoint(),
		TopicPaymentSuccess,
		ConsumerGroupShipment,
		seaman.OffsetResetToEarliest,
		func(ctx seaman.IContext) error {
			service := NewProductService(&ProductServiceOptions{
				Context: ctx,
				Config:  svc.config,
				Store: NewProductStore(&ProductStoreOptions{
					Context: ctx,
					Config:  svc.config,
				}),
			})
			return svc.handleCreateShipmentInfo(ctx, service)
		}))
}

func (svc *ProductMQ) handleCreateShipmentInfo(ctx seaman.IContext, productService IProductService) error {
	payload := &OrderPayload{}
	err := ctx.Transformer().TransformJSON(ctx.ReadInput(), payload, true)
	if err != nil {
		return ctx.NewError(err, errors.InternalServerError)
	}

	_, ierr := productService.Delete("")
	if ierr != nil {
		return ctx.NewError(ierr, ierr)
	}

	return nil
}

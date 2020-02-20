// Copyright © PushAndMotion 2018 All Rights Reserved.
// PAM Engine & Library is proprietary and confidential.
// Un-authorize using, editing, copying, adapting, distributing, of this file or some part of this file without
// the prior written consent of PushAndMotion, via any medium is strictly prohibited.
// If not expressively specify in the document, the authorisation to use this library will be granted per application.
// Any question regarding this copyright notice please contact contact@pushandmotion.com
// This copyright notice must be included in the header of every distribution of all the source code.

package ecommons

import (
	"fmt"

	"bitbucket.org/3dsinteractive/pam4/config"
	"bitbucket.org/3dsinteractive/pam4/f"
	"bitbucket.org/3dsinteractive/seaman"
)

// FAuthorizationHandler is authorize handler for frontend
func FAuthorizationHandler(ctx seaman.IContext) (bool, error) {
	accessToken := ctx.Header("Authorization")
	if len(accessToken) == 0 {
		return false, nil
	}

	cfg := config.NewConfig()
	idx := ctx.Indexer(cfg.IndexerConfig())

	query := f.RemoveTabAndNewLine(fmt.Sprintf(`{
		"query": {
		  "bool": {
			"filter": [
			  {
				"term": {
				  "tokens.token": "%s"
				}
			  },
			  {
				"term": {
				  "is_active": true
				}
			  }
			]
		  }
		}
	  }`, accessToken))

	items, _, err := idx.Query("members", query)
	if err != nil {
		return false, err
	}

	if len(items) == 0 {
		return false, nil
	}

	return true, nil
}

package core

import (
	"bitbucket.org/3dsinteractive/pam4/bconfig"
	"bitbucket.org/3dsinteractive/pam4/config"
	"bitbucket.org/3dsinteractive/seaman"
	"golang-guideline/models"
	"strings"
)

func GetPageOptions(params seaman.KeyValue, options *models.GetPageOptionsConfig) *models.PageOptions {
	defaultLimit := config.EComDefaultFrontendPageSize
	if options != nil && options.DefaultLimit > 0 {
		defaultLimit = options.DefaultLimit
	}

	pageLimit, _ := params.GetInt(config.ParamPageSizeQueryString.Name, defaultLimit)
	page, _ := params.GetInt(config.ParamPageQueryString.Name, 1)
	paramQ, _ := params.GetString(bconfig.ParamQueryString.Name, "")
	if len(strings.TrimSpace(paramQ)) == 0 {
		paramQ, _ = params.GetString(config.ParamKeywordQueryString.Name, "")
	}

	return &models.PageOptions{
		Q:     strings.TrimSpace(paramQ),
		Limit: pageLimit,
		Page:  page,
	}
}

func GetBeforeOptions(params seaman.KeyValue, options *models.GetPageOptionsConfig) *models.PageOptions {
	defaultLimit := config.EComDefaultFrontendPageSize
	if options != nil && options.DefaultLimit > 0 {
		defaultLimit = options.DefaultLimit
	}
	pageLimit, _ := params.GetInt(config.ParamPageSizeQueryString.Name, defaultLimit)
	cursor, _ := params.GetString(config.ParamPageQueryString.Name, "")
	paramQ, _ := params.GetString(bconfig.ParamQueryString.Name, "")
	if len(strings.TrimSpace(paramQ)) == 0 {
		paramQ, _ = params.GetString(config.ParamKeywordQueryString.Name, "")
	}

	return &models.PageOptions{
		Q:      strings.TrimSpace(paramQ),
		Limit:  pageLimit,
		Cursor: cursor,
		Page:   1,
	}
}

func GetAfterOptions(params seaman.KeyValue, options *models.GetPageOptionsConfig) *models.PageOptions {
	defaultLimit := config.EComDefaultFrontendPageSize
	if options != nil && options.DefaultLimit > 0 {
		defaultLimit = options.DefaultLimit
	}
	pageLimit, _ := params.GetInt(config.ParamPageSizeQueryString.Name, defaultLimit)
	cursor, _ := params.GetString(config.ParamPageQueryString.Name, "")
	paramQ, _ := params.GetString(bconfig.ParamQueryString.Name, "")
	if len(strings.TrimSpace(paramQ)) == 0 {
		paramQ, _ = params.GetString(config.ParamKeywordQueryString.Name, "")
	}

	return &models.PageOptions{
		Q:      strings.TrimSpace(paramQ),
		Limit:  pageLimit,
		Cursor: cursor,
		Page:   1,
	}
}

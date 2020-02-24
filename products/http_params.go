package products

import "bitbucket.org/3dsinteractive/seaman"

var paramPageQueryString = &seaman.AttributeParam{
	Name:       "page",
	Type:       seaman.AttributeTypeString,
	LocateIn:   seaman.ParamLocationQueryString,
	IsRequired: false,
	Validators: []*seaman.AttributeValidator{},
}

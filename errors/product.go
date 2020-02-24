package errors

import (
	"bitbucket.org/3dsinteractive/seaman"
	"net/http"
)

var (
	ProductImportFileTypeNotSupport = seaman.Error{
		Status:  http.StatusBadRequest,
		Code:    "FILE_TYPE_NOT_SUPPORT",
		Message: "please try .csv or .xlxs",
	}
)
